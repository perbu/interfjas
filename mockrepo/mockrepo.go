package mockrepo

import "github.com/perbu/interfjas/repo"

type mockRepository struct {
	state string
}

func (r mockRepository) GetRecord() repo.Record {
	return repo.Record{
		No:   666,
		Name: "Mocked Zonk",
	}
}

func NewRepository(state string) *mockRepository {
	return &mockRepository{
		state: state,
	}
}
