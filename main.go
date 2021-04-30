package main

import (
	"fmt"
)

// DB & ENTITY

// ------------------------------------------------------------------------------------------------------------
type (
	Student struct {
		ID			int
		Identifier	string
		Name		string
	}

	DB struct {
		database []Student
	}
)
// ------------------------------------------------------------------------------------------------------------

// CONTROLLER

// ------------------------------------------------------------------------------------------------------------
type StudentRepository interface {
	GetAll() []Student
	GetById(ID int) *Student
	Save(student *Student) *Student
	DeleteById(ID int) *Student
	UpdateById(ID int, newStudent *Student)	*Student
}

type mySqlRepo struct {
	db *DB
}

func newStudentRepository(db *DB) StudentRepository {
	return &mySqlRepo{db: db}
}

func (repo *mySqlRepo) GetAll() []Student {
	return repo.db.database
}

func (repo *mySqlRepo) GetById(ID int) *Student {
	for _, v := range repo.db.database {
		if v.ID == ID {
			return &v
		}
	}
	return nil
}

func (repo *mySqlRepo) Save(student *Student) *Student {
	id := len(repo.db.database)
	if id == 0 {
		student.ID = 0
	}else {
		student.ID = id
	}

	repo.db.database = append(repo.db.database, *student)
	return student
}

func (repo *mySqlRepo) DeleteById(ID int) *Student {
	student := repo.GetById(ID)
	temp := repo.db.database[:ID]
	repo.db.database =  append(temp, repo.db.database[ID+1:]...)
	return student
}

func (repo *mySqlRepo) UpdateById(ID int, newStudent *Student) *Student {
	var std Student

	for _, v := range repo.db.database {
		if v.ID == ID {
			std = v
		}
	}

	id := std.ID

	temp := repo.db.database[ID:]
	repo.db.database = append(temp, repo.db.database[ID+1:]...)

	newStudent.ID = id
	repo.Save(newStudent)
	return newStudent
}
// ------------------------------------------------------------------------------------------------------------

// SERVICE

// ------------------------------------------------------------------------------------------------------------
type StudentService interface {
	FindAllStudents() []Student
	FindStudentById(ID int) *Student
	CreateStudent(student *Student) *Student
	DeleteStudentById(ID int) *Student
	UpdateById(ID int, newStudent *Student) *Student
}

type studentService struct {
	repository StudentRepository
}

func newStudentService(repo StudentRepository) StudentService {
	return &studentService{repository: repo}
}

func (service *studentService) FindAllStudents() []Student {
	return service.repository.GetAll()
}

func (service *studentService) FindStudentById(ID int) *Student {
	return service.repository.GetById(ID)
}

func (service *studentService) CreateStudent(student *Student) *Student {
	return service.repository.Save(student)
}

func (service *studentService) DeleteStudentById(ID int) *Student {
	return service.repository.DeleteById(ID)
}

func (service *studentService) UpdateById(ID int, newStudent *Student) *Student {
	return service.repository.UpdateById(ID, newStudent)
}
// ------------------------------------------------------------------------------------------------------------

// CONTROLLER

// ------------------------------------------------------------------------------------------------------------
type StudentController interface {
	GetAllStudents()
	GetStudentById()
	SaveStudent()
	DeleteStudentById()
	UpdateStudentById()
}

type studentController struct {
	studentService StudentService
}

func newStudentController(service StudentService) StudentController {
	return &studentController{service}
}

func (s studentController) GetAllStudents() {
	fmt.Println(s.studentService.FindAllStudents())
}

func (s studentController) GetStudentById() {
	fmt.Print("Masukkan ID : ")
	var ID int
	fmt.Scan(&ID)
	fmt.Println(s.studentService.FindStudentById(ID))
}

func (s studentController) SaveStudent() {
	var stdIN Student
	var identifier,name string

	fmt.Print("IDENTIFIER : ")
	fmt.Scan(&identifier)

	fmt.Print("NAME : ")
	fmt.Scan(&name)

	stdIN = Student{Identifier: identifier, Name: name}
	result := s.studentService.CreateStudent(&stdIN)
	fmt.Println(result)
}

func (s studentController) DeleteStudentById() {
	fmt.Print("Masukkan ID : ")
	var ID int
	fmt.Scan(&ID)
	result := s.studentService.DeleteStudentById(ID)
	fmt.Println(result)
}

func (s studentController) UpdateStudentById() {
	fmt.Print("Masukkan ID : ")
	var ID int
	fmt.Scan(&ID)

	var identifier,name string

	fmt.Print("IDENTIFIER : ")
	fmt.Scan(&identifier)

	fmt.Print("NAME : ")
	fmt.Scan(&name)
	newStudent := Student{Identifier: identifier, Name: name}
	fmt.Println(s.studentService.UpdateById(ID, &newStudent))
}
// ------------------------------------------------------------------------------------------------------------

func main(){
	var db []Student
	database := DB{database: db}

	// INJECT
	repository := newStudentRepository(&database)
	service := newStudentService(repository)
	controller := newStudentController(service)

	controller.SaveStudent()
	controller.SaveStudent()
	controller.SaveStudent()

	controller.GetAllStudents()

	controller.GetStudentById()
	controller.DeleteStudentById()
	controller.GetAllStudents()

	controller.UpdateStudentById()
	controller.GetAllStudents()
}