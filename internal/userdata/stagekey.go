package userdata

type Stagekey int

// func (s *Stagekey) String() string {
// 	if s != nil {
// 		return string(*s)
// 	}
// 	return ""
// }

const (
	Stage_Start Stagekey = iota
)

func (c Stagekey) String() string {
	switch c {
	case Stage_Start:
		return "Стадия: старт"

	}
	return "unknown"
}
