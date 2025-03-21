package services_test

import (
	"os"
	"testing"

	"github.com/markitos/markitos-svc-boilerplate/internal/domain"
	internal_test "github.com/markitos/markitos-svc-boilerplate/testsuite/internal"
)

type MockSpyBoilerRepository struct {
	LastCreatedBoilerName *string
	LastDeleteBoilerId    *string
	LastOneBoilerId       *string
	LastUpdatedBoilerId   *string
	LastUpdatedBoilerName *string
}

func NewMockSpyBoilerRepository() *MockSpyBoilerRepository {
	return &MockSpyBoilerRepository{
		LastCreatedBoilerName: nil,
		LastDeleteBoilerId:    nil,
		LastOneBoilerId:       nil,
	}
}

func (m *MockSpyBoilerRepository) Create(boiler *domain.Boilerplate) error {
	m.LastCreatedBoilerName = &boiler.Name

	return nil
}

func (m *MockSpyBoilerRepository) CreateHaveBeenCalledWith(boilerName *string) bool {
	var result bool = m.LastCreatedBoilerName != nil && *m.LastCreatedBoilerName == *boilerName

	m.LastCreatedBoilerName = nil

	return result
}

func (m *MockSpyBoilerRepository) Delete(id *domain.BoilerplateId) error {
	value := id.Value()
	m.LastDeleteBoilerId = &value

	return nil
}

func (m *MockSpyBoilerRepository) DeleteHaveBeenCalledWith(boilerId *string) bool {
	var result bool = m.LastDeleteBoilerId != nil && *m.LastDeleteBoilerId == *boilerId

	m.LastDeleteBoilerId = nil

	return result
}

func (m *MockSpyBoilerRepository) Update(boiler *domain.Boilerplate) error {
	m.LastUpdatedBoilerId = &boiler.Id
	m.LastUpdatedBoilerName = &boiler.Name

	return nil
}

func (m *MockSpyBoilerRepository) UpdateHaveBeenCalledWith(id, name string) bool {
	var matchId bool = *m.LastUpdatedBoilerId == id
	var matchName bool = *m.LastUpdatedBoilerName == name
	var matchWithOneCall bool = m.LastOneBoilerId != nil && *m.LastOneBoilerId == id

	m.LastUpdatedBoilerId = nil
	m.LastUpdatedBoilerName = nil
	m.LastOneBoilerId = nil

	return matchId && matchName && matchWithOneCall
}

func (m *MockSpyBoilerRepository) One(id *domain.BoilerplateId) (*domain.Boilerplate, error) {
	value := id.Value()
	m.LastOneBoilerId = &value

	return internal_test.NewRandomBoilerplateWithCustomId(id), nil
}

func (m *MockSpyBoilerRepository) OneHaveBeenCalledWith(boilerId *string) bool {
	var result bool = m.LastOneBoilerId != nil && *m.LastOneBoilerId == *boilerId

	m.LastOneBoilerId = nil

	return result
}

var repository *MockSpyBoilerRepository

func TestMain(m *testing.M) {
	repository = NewMockSpyBoilerRepository()

	os.Exit(m.Run())
}
