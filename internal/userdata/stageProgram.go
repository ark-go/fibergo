package userdata

type ProgName string

// Step Programm/Step
//type Stage map[ProgName]string

type NextStep struct {
	Program string
	Step    string
}

func (u *User) ChangeProgram(str string) string {
	u.UserData.NextStep.Program = str
	return str
}
func (u *User) ChangeStep(str string) string {
	u.UserData.NextStep.Step = str
	return str
}
func (u *User) GetStep() (prog, step string) {
	prog = u.UserData.NextStep.Program
	step = u.UserData.NextStep.Step

	return
}
