package model

import (
	"fmt"

	"github.com/lukpank/go-glpk/glpk"
)

type Scenario struct {
	Employees []string                      `json:"employees"`
	Shifts    []string                      `json:"shifts"`
	Cost      map[string]map[string]float64 `json:"cost"`
	Demand    map[string]int                `json:"demand"`
}

type ScheduleResult struct {
	Assignments map[string]string `json:"assignments"`
	TotalCost   float64           `json:"total_cost"`
}

func SolveShiftSchedule(s Scenario) (ScheduleResult, error) {
	prob := glpk.New()
	defer prob.Delete()

	// Log demands
	fmt.Println("Dataset Demands:")
	for _, sh := range s.Shifts {
		fmt.Printf("  %s: %d\n", sh, s.Demand[sh])
	}

	varIndex := map[string]int{}
	varCount := 0

	// 1) Add binary variables X_emp_shift
	for _, emp := range s.Employees {
		for _, sh := range s.Shifts {
			c, ok := s.Cost[emp][sh]
			if !ok || c <= 0 {
				fmt.Printf("Skipping invalid shift: %s → %s\n", emp, sh)
				continue
			}
			varCount++
			name := fmt.Sprintf("X_%s_%s", emp, sh)
			prob.AddCols(1)
			prob.SetColName(varCount, name)
			prob.SetColKind(varCount, glpk.VarType(glpk.BV))
			prob.SetColBnds(varCount, glpk.BndsType(glpk.DB), 0.0, 1.0)
			prob.SetObjCoef(varCount, c)
			varIndex[name] = varCount
			fmt.Printf("Added variable: %s (col %d, cost %.0f)\n", name, varCount, c)
		}
	}
	fmt.Printf("Total variables added: %d\n", varCount)
	prob.SetObjDir(glpk.ObjDir(glpk.MIN))

	// 2) Constraint: Each shift must meet demand
	for _, sh := range s.Shifts {
		var idx []int32
		var val []float64
		for _, emp := range s.Employees {
			if _, ok := s.Cost[emp][sh]; !ok || s.Cost[emp][sh] <= 0 {
				continue
			}
			name := fmt.Sprintf("X_%s_%s", emp, sh)
			idx = append(idx, int32(varIndex[name]))
			val = append(val, 1)
		}
		prob.AddRows(1)
		row := prob.NumRows()
		prob.SetRowName(row, fmt.Sprintf("Shift_%s", sh))
		prob.SetRowBnds(row, glpk.BndsType(glpk.FX), float64(s.Demand[sh]), float64(s.Demand[sh]))
		prob.SetMatRow(row, idx, val)
	}

	// 3) Constraint: Each employee can have at most one shift
	for _, emp := range s.Employees {
		var idx []int32
		var val []float64
		for _, sh := range s.Shifts {
			if _, ok := s.Cost[emp][sh]; !ok || s.Cost[emp][sh] <= 0 {
				continue
			}
			name := fmt.Sprintf("X_%s_%s", emp, sh)
			idx = append(idx, int32(varIndex[name]))
			val = append(val, 1)
		}
		if len(idx) == 0 {
			continue
		}
		prob.AddRows(1)
		row := prob.NumRows()
		prob.SetRowName(row, fmt.Sprintf("Employee_%s", emp))
		prob.SetRowBnds(row, glpk.BndsType(glpk.UP), 0.0, 1.0)
		prob.SetMatRow(row, idx, val)
	}

	// 4) Write LP model
	prob.WriteLP(nil, "model.lp")

	// 5) LP-relaxation
	if err := prob.Simplex(nil); err != nil {
		return ScheduleResult{}, fmt.Errorf("Simplex failed: %w", err)
	}
	if prob.Status() != glpk.OPT {
		return ScheduleResult{}, fmt.Errorf("Simplex not optimal (status %d)", prob.Status())
	}

	// 6) MIP solve
	iocp := glpk.NewIocp()
	iocp.SetPresolve(false)
	if err := prob.Intopt(iocp); err != nil {
		return ScheduleResult{}, fmt.Errorf("MIP failed: %w", err)
	}
	if prob.MipStatus() != glpk.OPT {
		return ScheduleResult{}, fmt.Errorf("MIP not optimal (status %d)", prob.MipStatus())
	}

	// 7) Parse MIP solution
	assignments := make(map[string]string)
	shiftCounts := make(map[string]int)
	var total float64

	for _, emp := range s.Employees {
		assigned := false
		for _, sh := range s.Shifts {
			colName := fmt.Sprintf("X_%s_%s", emp, sh)
			col, ok := varIndex[colName]
			if !ok {
				continue
			}
			val := prob.MipColVal(col)
			fmt.Printf("MIP Variable %s (col %d): value = %.2f\n", colName, col, val)
			if val >= 0.99 && !assigned {
				assignments[emp] = sh
				shiftCounts[sh]++
				total += s.Cost[emp][sh]
				assigned = true
			}
		}
	}

	// 8) Fallback: Assign unassigned employees if understaffed
	for _, sh := range s.Shifts {
		if shiftCounts[sh] < s.Demand[sh] {
			for _, emp := range s.Employees {
				if _, assigned := assignments[emp]; !assigned {
					if c, ok := s.Cost[emp][sh]; ok && c > 0 {
						assignments[emp] = sh
						shiftCounts[sh]++
						total += c
						fmt.Printf("Fallback: Assigned %s → %s\n", emp, sh)
						if shiftCounts[sh] >= s.Demand[sh] {
							break
						}
					}
				}
			}
		}
	}

	// 9) Log assignments and counts
	fmt.Println("MIP Assignments:")
	for emp, sh := range assignments {
		fmt.Printf("  %s → %s\n", emp, sh)
	}
	fmt.Println("Jumlah pegawai per shift (calculated):")
	for _, sh := range s.Shifts {
		count := 0
		for _, assignedShift := range assignments {
			if assignedShift == sh {
				count++
			}
		}
		fmt.Printf("- %s: %d orang (expected %d)\n", sh, count, s.Demand[sh])
		shiftCounts[sh] = count
	}

	// 10) Validate solution
	for _, sh := range s.Shifts {
		if shiftCounts[sh] != s.Demand[sh] {
			return ScheduleResult{}, fmt.Errorf(
				"MIP shift %s has %d employees, required %d",
				sh, shiftCounts[sh], s.Demand[sh],
			)
		}
	}

	return ScheduleResult{
		Assignments: assignments,
		TotalCost:   total,
	}, nil
}
