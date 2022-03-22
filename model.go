package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

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

type Data struct {
	ID      string `json:"ID" validate:"required"`
	Message string `json:"Message" validate:"required"`
}

func (d Data) Json() string {
	out, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(out)
}

func (d Data) String() string {
	out := fmt.Sprintf("{ID: %s, Message: %s}", d.ID, d.Message)
	return out
}

type AllData map[string]Data

func (a AllData) String() string {
	returnData := ""

	for _, x := range serviceData {
		returnData += "{" + x.Json() + "}"
	}

	return "[" + returnData + "]"
}

var serviceData AllData

type Logic struct {
	serviceData AllData
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

func (a AllData) search(id string) (Data, bool) {
	d, found := a[id]
	return d, found
}

func (l *Logic) Create(newData Data) error {
	l.m.Lock()
	defer l.m.Unlock()

	if _, found := l.serviceData.search(newData.ID); found {
		return ErrDataConflict
	}

	l.serviceData[newData.ID] = newData
	return nil
}

func (l *Logic) Read(id string) (Data, bool) {
	l.m.Lock()
	defer l.m.Unlock()

	return l.serviceData.search(id)
}

func ReadAll() AllData {
	l.m.Lock()
	defer l.m.Unlock()

	// returning a copy
	out := AllData{}
	for k, v := range l.serviceData {
		out[k] = v
	}

	return out
}

func (l *Logic) Update(input Data) (Data, error) {
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

func Delete(id string) error {
	l.m.Lock()
	defer l.m.Unlock()

	if _, found := l.serviceData[id]; !found {
		return ErrDataNotFound
	}

	delete(l.serviceData, id)
	return nil
}

//func MockData(overwrite bool, data AllData) error {
//
//	if overwrite {
//		serviceData = data
//		return nil
//	} else {
//		for i, singleData := range data {
//
//		}
//	}
//
//}
