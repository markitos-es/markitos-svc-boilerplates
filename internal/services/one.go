package services

import "github.com/markitos-es/markitos-svc-boilerplates/internal/domain"

type BoilerplateOneRequest struct {
	Id string `json:"id"`
}

type BoilerplateOneResponse struct {
	Data *domain.Boilerplate `json:"data"`
}

type BoilerplateOneService struct {
	Repository domain.BoilerplateRepository
}

func NewBoilerplateOneService(repository domain.BoilerplateRepository) BoilerplateOneService {
	return BoilerplateOneService{
		Repository: repository,
	}
}

func (s BoilerplateOneService) Do(request BoilerplateOneRequest) (*BoilerplateOneResponse, error) {
	boilerId, err := domain.NewBoilerplateId(request.Id)
	if err != nil {
		return nil, err
	}

	boilerplate, err := s.Repository.One(boilerId)
	if err != nil {
		return nil, err
	}

	return &BoilerplateOneResponse{boilerplate}, nil
}
