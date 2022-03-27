package main

import (
	"encoding/json"
	"errors"
	Log "github.com/sirupsen/logrus"
	"jimmyray.io/data-api/utils"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/go-playground/validator/v10"
)

const (
	HttpReqReadErr    string = "HTTP_REQ_READ_ERR"
	IncorrectInputErr string = "Please check submission: {\"ID\":\"<ID_VALUE>\",\"Message\":\"<MESSAGE_VALUE>\"}"
	JsonEncodeErr     string = "JSON_ENCODE_ERR"
	DataConflictErr   string = "NOOP_DATA_CONFLICT_ERR"
	DataNotFoundErr   string = "NOOP_DATA_NOT_FOUND_ERR"
	ValidationErr     string = "VALIDATION_ERR"
	InternalServerErr string = "INTERNAL_SERVER_ERR"
	MockDataErr       string = "Mock_Data_Err"
)

var (
	ErrDataNotFound   = errors.New(DataNotFoundErr)
	ErrDataConflict   = errors.New(DataConflictErr)
	ErrInternalServer = errors.New(InternalServerErr)
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()
}

var serviceId = uuid.New()

func GetServiceId() string {
	return serviceId.String()
}

//Employee
type Department struct {
	ID        string `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	ManagerID string `json:"mgrId" validate:"required"`
}

type Address struct {
	Street  string `json:"street" validate:"required"`
	City    string `json:"city" validate:"required"`
	County  string `json:"county" validate:"required"`
	State   string `json:"state" validate:"required"`
	Zipcode string `json:"zipcode" validate:"required"`
}

type Employee struct {
	ID       string     `json:"id" validate:"required"`
	FName    string     `json:"fname" validate:"required"`
	LName    string     `json:"lname" validate:"required"`
	Sex      string     `json:"sex" validate:"required"`
	DOB      time.Time  `json:"dob" validate:"required"`
	HireDate time.Time  `json:"hireDate" validate:"required"`
	Position string     `json:"position" validate:"required"`
	Salary   uint64     `json:"salary" validate:"required"`
	Dept     Department `json:"dept" validate:"required"`
	Address  Address    `json:"address" validate:"required"`
}

func (e Employee) Json() string {
	out, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	return string(out)
}

//type Data struct {
//	ID      string `json:"ID" validate:"required"`
//	Message string `json:"Message" validate:"required"`
//}

//func (d Data) Json() string {
//	out, err := json.Marshal(d)
//	if err != nil {
//		panic(err)
//	}
//
//	return string(out)
//}
//
//func (d Data) String() string {
//	out := fmt.Sprintf("{ID: %s, Message: %s}", d.ID, d.Message)
//	return out
//}

//type AllData map[string]Data

//func (a AllData) String() string {
//	returnData := ""
//
//	for _, x := range serviceData {
//		returnData += "{" + x.Json() + "}"
//	}
//
//	return "[" + returnData + "]"
//}

type employees map[string]Employee

func (e employees) Json() string {
	out, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	return string(out)
}

//var serviceData AllData
//
//type Logic struct {
//	serviceData AllData
//	m           sync.Mutex
//}

//var serviceData employees

type Logic struct {
	serviceData employees
	m           sync.Mutex
}

type Controller struct {
	l        *Logic
	validate *validator.Validate
}

type info struct {
	NAME string `json:"service-name"`
	ID   string `json:"service-id"`
}

func (i info) String() string {
	out, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	return string(out)
}

var ServiceInfo = info{}

type InfoController struct {
	ServiceInfo info
}

var (
	l  Logic
	ic InfoController
	c  Controller
)

//func (a AllData) search(id string) (Data, bool) {
//	d, found := a[id]
//	return d, found
//}

func (e employees) search(id string) (Employee, bool) {
	d, found := e[id]
	return d, found
}

//func (l *Logic) Create(newData Data) error {
//	l.m.Lock()
//	defer l.m.Unlock()
//
//	if _, found := l.serviceData.search(newData.ID); found {
//		return ErrDataConflict
//	}
//
//	l.serviceData[newData.ID] = newData
//	return nil
//}

func (l *Logic) Create(newData Employee) error {
	l.m.Lock()
	defer l.m.Unlock()

	if _, found := l.serviceData.search(newData.ID); found {
		return ErrDataConflict
	}

	l.serviceData[newData.ID] = newData
	return nil
}

//func (l *Logic) Read(id string) (Data, bool) {
//	l.m.Lock()
//	defer l.m.Unlock()
//
//	return l.serviceData.search(id)
//}

func (l *Logic) Read(id string) (Employee, bool) {
	l.m.Lock()
	defer l.m.Unlock()

	return l.serviceData.search(id)
}

//func ReadAll() AllData {
//	l.m.Lock()
//	defer l.m.Unlock()
//
//	// returning a copy
//	out := AllData{}
//	for k, v := range l.serviceData {
//		out[k] = v
//	}
//
//	return out
//}

func ReadAll() employees {
	l.m.Lock()
	defer l.m.Unlock()

	// returning a copy
	out := employees{}
	for k, v := range l.serviceData {
		out[k] = v
	}

	return out
}

//func (l *Logic) Update(input Data) (Data, error) {
//	l.m.Lock()
//	defer l.m.Unlock()
//
//	foundData, found := l.serviceData[input.ID]
//	if !found {
//		return foundData, ErrDataNotFound
//	}
//	if foundData == input {
//		return foundData, ErrDataConflict
//	}
//	l.serviceData[input.ID] = input
//	return l.serviceData[input.ID], nil
//}

func (l *Logic) Update(input Employee) (Employee, error) {
	l.m.Lock()
	defer l.m.Unlock()

	foundData, found := l.serviceData[input.ID]
	if !found {
		return foundData, ErrDataNotFound
	}
	if foundData == input {
		return foundData, ErrDataConflict
	}
	l.serviceData[input.ID] = input
	return l.serviceData[input.ID], nil
}

//func Delete(id string) error {
//	l.m.Lock()
//	defer l.m.Unlock()
//
//	if _, found := l.serviceData[id]; !found {
//		return ErrDataNotFound
//	}
//
//	delete(l.serviceData, id)
//	return nil
//}

func Delete(id string) error {
	l.m.Lock()
	defer l.m.Unlock()

	if _, found := l.serviceData[id]; !found {
		return ErrDataNotFound
	}

	delete(l.serviceData, id)
	return nil
}

func loadMockData() error {
	e := employees{}
	err := json.Unmarshal([]byte(mockData), &e)

	if err != nil {
		utils.Logger.WithFields(Log.Fields{"error": err.Error()}).Debug("")
	}
	utils.Logger.WithFields(Log.Fields{"length": len(e)}).Debug("Employee map length")

	if err == nil {
		l.m.Lock()
		defer l.m.Unlock()

		for k, v := range e {
			l.serviceData[k] = v
		}
	}

	utils.Logger.WithFields(Log.Fields{"length": len(l.serviceData)}).Debug("ServiceData map length")

	return err
}
