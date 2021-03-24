package main

import (
	"fmt"
	"github.com/perbu/interfjas/mockrepo"
	"github.com/perbu/interfjas/repo"
)

func (s *service) getRepo() *repo.Repository {
	if s.serviceRepo != nil {
		return &s.serviceRepo
	}
	s.serviceRepo = repo.NewRepository("the real shit")
	return &s.serviceRepo
}

func (s *service) getMockedRepo() *repo.Repository {
	if s.serviceRepo != nil {
		return &s.serviceRepo
	}
	s.serviceRepo = mockrepo.NewRepository("the mocked shit")
	return &s.serviceRepo

}

type service struct {
	serviceRepo repo.Repository
}

func main() {
	var s service
	r := s.getRepo()
	fmt.Println("We're up and running.")
	fmt.Println("repo:", *r)
	fmt.Println("Record: ", (*r).GetRecord())

	var ms service
	mr := ms.getMockedRepo()
	fmt.Println("Mocking....")
	fmt.Println("repo:", *mr)
	fmt.Println("Record: ", (*mr).GetRecord())

}
