package chineseReminderTheory

type ReminderClass struct {
	modulus  float64
	reminder float64
}

func CreateReminderClass(modulus, reminder float64) *ReminderClass {
	return &ReminderClass{modulus: modulus, reminder: reminder}
}

type ReminderClasses []*ReminderClass

func (r *ReminderClasses) Initialize() {
	*r = make([]*ReminderClass, 0)
}

func (r *ReminderClasses) Add(reminderClass *ReminderClass) {
	*r = append(*r, reminderClass)
}

// multiplyAllModulus multiplies all ReminderClass.modulus together.
func (r *ReminderClasses) multiplyAllModulus() float64 {
	allModulus := 1.0
	for _, r := range *r {
		allModulus *= r.modulus
	}
	return allModulus
}
