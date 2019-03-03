package resources

type Building struct {
	Id string			`json:"id"`
	Name string			`json:"name"`
	NumFloors int		`json:"numFloors"`
}

func (b *Building) Create() string {
	return "1"
}

func (b *Building) Get() *Building {
	building := Building{Id: "1", Name: "Empire State", NumFloors: 100}
	return &building
}
