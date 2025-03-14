package domain

type BoilerplateRepository interface {
	Create(boiler *Boilerplate) error
	Delete(id *BoilerplateId) error
	One(id *BoilerplateId) (*Boilerplate, error)
	Update(boilerplate *Boilerplate) error
}
