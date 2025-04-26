package main

import (
	"fmt"
	"go-model-employee-shifting/model"
)

func LoadScenario() model.Scenario {
	return model.Scenario{
		Employees: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"},
		Shifts:    []string{"Pagi", "Siang", "Malam"},
		Cost: map[string]map[string]float64{
			"A": {"Pagi": 50000, "Siang": 60000},
			"B": {"Siang": 45000, "Malam": 55000},
			"C": {"Pagi": 55000, "Malam": 50000},
			"D": {"Pagi": 58000, "Siang": 48000, "Malam": 60000},
			"E": {"Pagi": 60000, "Siang": 70000},
			"F": {"Siang": 65000, "Malam": 55000},
			"G": {"Pagi": 59000, "Siang": 61000, "Malam": 59000},
			"H": {"Pagi": 56000, "Siang": 60000, "Malam": 60000},
			"I": {"Pagi": 58000, "Siang": 61000, "Malam": 60000},
			"J": {"Pagi": 50000, "Siang": 46000, "Malam": 58000},
			"K": {"Pagi": 52000, "Siang": 62000, "Malam": 57000},
			"L": {"Pagi": 53000, "Siang": 64000, "Malam": 56000},
		},
		Demand: map[string]int{
			"Pagi":  3,
			"Siang": 4,
			"Malam": 3,
		},
	}
}

func PrintResult(res model.ScheduleResult) {
	fmt.Println("=== Jadwal Pegawai ===")
	for emp, shift := range res.Assignments {
		fmt.Printf("  Pegawai %s → %s\n", emp, shift)
	}
	fmt.Printf("Total Biaya: Rp %.0f\n", res.TotalCost)
}

func main() {
	// Load scenario
	scenario := LoadScenario()

	// Build and solve model
	result, err := model.SolveShiftSchedule(scenario)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Output results
	PrintResult(result)
}
