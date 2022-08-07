package server

type Per struct {
	name string
	age  int
}

func (p Per) Name() string {
	return p.name
}

func (p *Per) SetName(NewName string) {
	p.name = NewName
}
