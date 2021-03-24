package repo

type Record struct {
	No   int
	Name string
}

type Repository interface {
	GetRecord() Record
}

type repository struct {
	state string
}

func (r *repository) GetRecord() Record {
	return Record{
		No:   1,
		Name: "Zonk",
	}

}

func (r *repository) SetState(s string) {
	r.state = s
}

func NewRepository(state string) *repository {
	return &repository{
		state: state,
	}
}
