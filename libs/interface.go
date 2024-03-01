package libs

type Personer interface {
	GetName() string
	GetAge() int
}

type Student struct {
	Name string
	Age  int
}

func (s Student) GetName() string {
	return s.Name
}

func (s Student) GetAge() int {
	return s.Age
}

type Adulter interface {
	Personer
	GetCareer() string
}

type Man struct {
	Name   string
	Age    int
	Career string
}

func (m Man) GetName() string {
	return m.Name
}

func (m Man) GetAge() int {
	return m.Age
}

func (m Man) GetCareer() string {
	return m.Career
}
