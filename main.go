package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

type (
	Person struct {
		ID	 string
		Name string
	}

	PersonRepository struct {
		person []Person
	}

	PersonService struct {
		PersonRepository PersonRepository
	}

	PersonController struct {
		PersonService PersonService
	}

	personRepository interface {
		GetPersonByName(name string) (*Person, error)
		Save(person *Person) bool
	}

	personService interface {
		FindPersonByName(name string) (*Person, error)
		SavePerson(person *Person) bool
	}

	personController interface {
		FindPersonByName()
		SavePerson()
	}
)

// Dummy
func NewDummyPerson() *[]Person {
	return &[]Person{
		{"1","Sam"},
		{"2","Dev"},
		{"3","sammidev"},
	}
}

// Provide
func ProvidePersonRepository(person *[]Person) PersonRepository {
	return PersonRepository{person: *person}
}
func ProvidePersonService(p PersonRepository) PersonService {
	return PersonService{PersonRepository: p}
}
func ProvidePersonController(s PersonService) PersonController {
	return PersonController{PersonService: s}
}

// Impl
func repositoryImpl() personRepository{
	return &PersonRepository{}
}
func serviceImpl() personService{
	return &PersonService{}
}
func controllerImpl() personController{
	return &PersonController{}
}

// Repository
func (r *PersonRepository) GetPersonByName(name string)  (*Person, error){
	log.Println("REPOSITORY PROCESSING ... ")
	time.Sleep(1 * time.Second)
	for _, v := range r.person {
		if v.Name == name {
			return &v, nil
		}
	}
	return &Person{}, errors.New("PERSON TIDAK DITEMUKAN")
}
func (r *PersonRepository) Save(person *Person) bool {
	log.Println("REPOSITORY PROCESSING ... ")
	time.Sleep(1 * time.Second)
	first := len(r.person)
	r.person = append(r.person, *person)
	second := len(r.person)
	if second > first {
		return true
	}
	return false
}

// Service
func (s *PersonService) FindPersonByName(name string) (*Person, error) {
	log.Println("SERVICE LAYER PROCESSING ... ")
	time.Sleep(1 * time.Second)
	result, err :=s.PersonRepository.GetPersonByName(name)
	if err != nil {
		return &Person{}, err
	}
	return result, nil
}
func (s *PersonService) SavePerson(person *Person) bool {
	log.Println("SERVICE LAYER PROCESSING ... ")
	time.Sleep(1 * time.Second)
	return s.PersonRepository.Save(person)
}

// Controller
func (p *PersonController) FindPersonByName() {
	log.Println("CONTROLLER PROCESSING ... ")
	time.Sleep(1 * time.Second)

	var name string
	fmt.Print("masukkan nama yang ingin anda cari : ")
	_, _ = fmt.Scanln(&name)
	result,err := p.PersonService.FindPersonByName(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func (p *PersonController) SavePerson() {
	log.Println("CONTROLLER PROCESSING ... ")
	time.Sleep(1 * time.Second)

	var name string
	fmt.Print("masukkan nama : ")
	_, _ = fmt.Scanln(&name)

	var person = Person{
		ID: strconv.Itoa(len(p.PersonService.PersonRepository.person) + 1),
		Name: name,
	}

	result := p.PersonService.SavePerson(&person)
	fmt.Println(result)
}

// Main
func main(){
	fmt.Println("WRAP")

	person := NewDummyPerson()
	fmt.Println("db\t ->\t ", person)

	repo := ProvidePersonRepository(person)
	fmt.Println("repo\t ->\t ", repo)

	service := ProvidePersonService(repo)
	fmt.Println("service\t ->\t ",service)

	controller := ProvidePersonController(service)
	fmt.Println("controller ->\t ",  controller)
}