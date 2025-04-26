\* Problem: Unknown *\

Minimize
 obj: + 50000 X_A_Pagi + 60000 X_A_Siang + 45000 X_B_Siang
 + 55000 X_B_Malam + 55000 X_C_Pagi + 50000 X_C_Malam + 58000 X_D_Pagi
 + 48000 X_D_Siang + 60000 X_D_Malam + 60000 X_E_Pagi + 70000 X_E_Siang
 + 65000 X_F_Siang + 55000 X_F_Malam + 59000 X_G_Pagi + 61000 X_G_Siang
 + 59000 X_G_Malam + 56000 X_H_Pagi + 60000 X_H_Siang + 60000 X_H_Malam
 + 58000 X_I_Pagi + 61000 X_I_Siang + 60000 X_I_Malam + 50000 X_J_Pagi
 + 46000 X_J_Siang + 58000 X_J_Malam + 52000 X_K_Pagi + 62000 X_K_Siang
 + 57000 X_K_Malam + 53000 X_L_Pagi + 64000 X_L_Siang + 56000 X_L_Malam

Subject To
 Shift_Pagi: + X_L_Pagi + X_K_Pagi + X_J_Pagi + X_I_Pagi + X_H_Pagi
 + X_G_Pagi + X_E_Pagi + X_D_Pagi + X_C_Pagi = 3
 Shift_Siang: + X_L_Siang + X_K_Siang + X_J_Siang + X_I_Siang
 + X_H_Siang + X_G_Siang + X_F_Siang + X_E_Siang + X_D_Siang + X_B_Siang
 = 4
 Shift_Malam: + X_L_Malam + X_K_Malam + X_J_Malam + X_I_Malam
 + X_H_Malam + X_G_Malam + X_F_Malam + X_D_Malam + X_C_Malam = 3
 Employee_A: + X_A_Siang <= 1
 Employee_B: + X_B_Malam <= 1
 Employee_C: + X_C_Malam <= 1
 Employee_D: + X_D_Malam + X_D_Siang <= 1
 Employee_E: + X_E_Siang <= 1
 Employee_F: + X_F_Malam <= 1
 Employee_G: + X_G_Malam + X_G_Siang <= 1
 Employee_H: + X_H_Malam + X_H_Siang <= 1
 Employee_I: + X_I_Malam + X_I_Siang <= 1
 Employee_J: + X_J_Malam + X_J_Siang <= 1
 Employee_K: + X_K_Malam + X_K_Siang <= 1
 Employee_L: + X_L_Malam + X_L_Siang <= 1

Bounds
 0 <= X_A_Pagi <= 1
 0 <= X_A_Siang <= 1
 0 <= X_B_Siang <= 1
 0 <= X_B_Malam <= 1
 0 <= X_C_Pagi <= 1
 0 <= X_C_Malam <= 1
 0 <= X_D_Pagi <= 1
 0 <= X_D_Siang <= 1
 0 <= X_D_Malam <= 1
 0 <= X_E_Pagi <= 1
 0 <= X_E_Siang <= 1
 0 <= X_F_Siang <= 1
 0 <= X_F_Malam <= 1
 0 <= X_G_Pagi <= 1
 0 <= X_G_Siang <= 1
 0 <= X_G_Malam <= 1
 0 <= X_H_Pagi <= 1
 0 <= X_H_Siang <= 1
 0 <= X_H_Malam <= 1
 0 <= X_I_Pagi <= 1
 0 <= X_I_Siang <= 1
 0 <= X_I_Malam <= 1
 0 <= X_J_Pagi <= 1
 0 <= X_J_Siang <= 1
 0 <= X_J_Malam <= 1
 0 <= X_K_Pagi <= 1
 0 <= X_K_Siang <= 1
 0 <= X_K_Malam <= 1
 0 <= X_L_Pagi <= 1
 0 <= X_L_Siang <= 1
 0 <= X_L_Malam <= 1

Generals
 X_A_Pagi
 X_A_Siang
 X_B_Siang
 X_B_Malam
 X_C_Pagi
 X_C_Malam
 X_D_Pagi
 X_D_Siang
 X_D_Malam
 X_E_Pagi
 X_E_Siang
 X_F_Siang
 X_F_Malam
 X_G_Pagi
 X_G_Siang
 X_G_Malam
 X_H_Pagi
 X_H_Siang
 X_H_Malam
 X_I_Pagi
 X_I_Siang
 X_I_Malam
 X_J_Pagi
 X_J_Siang
 X_J_Malam
 X_K_Pagi
 X_K_Siang
 X_K_Malam
 X_L_Pagi
 X_L_Siang
 X_L_Malam

End
