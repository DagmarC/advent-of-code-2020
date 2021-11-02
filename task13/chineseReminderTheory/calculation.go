package chineseReminderTheory

import (
	"fmt"
	"math"
)

// CalculateChineseReminderTheory In its basic form, the Chinese remainder theorem will determine a number t that,
// when divided by some given divisors, leaves given remainders.
// Example:
// x ≡ 2 (mod 3),
//x ≡ 3 (mod 5),
//x ≡ 2 (mod 7)?
//______________________
//κ1 = 70 ≡ 0 (mod 5 · 7) ∧ κ1 = 70 ≡ 1 (mod 3),
//κ2 = 21 ≡ 0 (mod 3 · 7) ∧ κ2 = 21 ≡ 1 (mod 5),
//κ3 = 15 ≡ 0 (mod 3 · 5) ∧ κ3 = 15 ≡ 1 (mod 7).
// ______________________________________________
// Minimal value x jis given by congruency class modulo
//3 · 5 · 7 = 105, so
//x = 233 mod 105 = 23
func (r *ReminderClasses) CalculateChineseReminderTheory() float64 {

	x := 0.0
	allModulus := r.multiplyAllModulus()

	for _, reminderClass := range *r {
		zeroReminderClass := CreateReminderClass(allModulus/reminderClass.modulus, 0)

		oneReminderClass := CreateReminderClass(reminderClass.modulus, 1)

		commonReminder := findCommonReminderClass(oneReminderClass, zeroReminderClass)

		// xˆ = 2 · 70 + 3 · 21 + 2 · 15 = 233.
		x += reminderClass.reminder * commonReminder
	}
	return math.Mod(x, allModulus)
}

func findCommonReminderClass(oneReminderClass, zeroReminderClass *ReminderClass) float64 {
	// κ1 = 70 ≡ 0 (mod 5 · 7) ∧ κ1 = 70 ≡ 1 (mod 3),
	for k := zeroReminderClass.modulus; ; k += zeroReminderClass.modulus {
		if math.Mod(k, oneReminderClass.modulus) == oneReminderClass.reminder {
			return k
		}
	}
}

func Test() {
	var reminderClasses ReminderClasses
	reminderClasses.Initialize()

	reminderClasses.Add(CreateReminderClass(2, 1))
	reminderClasses.Add(CreateReminderClass(3, 2))
	reminderClasses.Add(CreateReminderClass(5, 4))
	reminderClasses.Add(CreateReminderClass(7, 0))

	fmt.Println(reminderClasses.CalculateChineseReminderTheory())

}
