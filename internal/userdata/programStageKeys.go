package userdata

// ключ для стадии
type Stagekey int

// ключ для программы
type Program int

// ключ для текстогово соответсвия связки ChatID:UserID
// определяется при каждом запросе поьзователя
type ChatUserStr string

// структура для Программы и стадии - швг в в этой программе
// нам необходимо хранить шаг пользователя, в базе данных
// для того чтобы можно было продолжить работу с пользователем если он ....
type StepUser struct {
	// стадия программы type Stagekey
	Stagekey Stagekey
	//  програма  type Program
	Program Program
}

// карта для хранения Программ в зависимости от ключа ChatUserStr
// для каждого чата и для каждого узера храним его шаг по конкретной программе
type MapStepUser map[ChatUserStr]StepUser

// используются для типа Stagekey в описаниях  StepUser
const (
	Stage_Start Stagekey = iota
)

// получить строковое представление значения Stagekey
func (c Stagekey) String() string {
	switch c {
	case Stage_Start:
		return "Стадия: старт"

	}
	return "unknown"
}

// используются для типа Program в описаниях  StepUser
const (
	Programm_Start Program = iota
	Programm_Volk
)

// получить стоковое представление для типа Program
func (p Program) String() string {
	switch p {
	case Programm_Start:
		return "Программа: старт"
	case Programm_Volk:
		return "Программа: Про волка"
	}
	return "Программа: unknown"
}
