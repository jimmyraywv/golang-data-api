package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	Log "github.com/sirupsen/logrus"
	"jimmyray.io/data-api/utils"

	"github.com/gorilla/mux"
)

func (ic InfoController) healthCheck(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "OK")
}

func (ic InfoController) getServiceInfo(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, ServiceInfo.String())
}

func (c *Controller) createData(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	var newData Employee

	err := json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		errorData := utils.ErrorLog{Skip: 1, Event: HttpReqReadErr, Message: err.Error()}
		utils.LogErrors(errorData)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(IncorrectInputErr))
		return
	}

	err = c.validate.Struct(newData)
	if err != nil {
		errorData := utils.ErrorLog{Skip: 1, Event: HttpReqReadErr, Message: err.Error(), ErrorData: string(newData.Json())}
		utils.LogErrors(errorData)

		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, IncorrectInputErr)
		return
	}

	err = c.l.Create(newData)

	if err != nil {
		if errors.Is(err, ErrDataConflict) {
			w.WriteHeader(http.StatusConflict)
			_, _ = fmt.Fprint(w, DataConflictErr)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(w, ErrInternalServer)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *Controller) getData(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	foundData, found := c.l.Read(id)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprint(w, DataNotFoundErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(foundData)
	if err != nil {
		errorData := utils.ErrorLog{Skip: 1, Event: JsonEncodeErr, Message: err.Error(), ErrorData: string(foundData.Json())}
		utils.LogErrors(errorData)
	}
}

func (c Controller) getAllData(w http.ResponseWriter, r *http.Request) {
	data := ReadAll()
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		errorData := utils.ErrorLog{Skip: 1, Event: JsonEncodeErr, Message: err.Error(), ErrorData: string(data.Json())}
		utils.LogErrors(errorData)
	}
}

func (c *Controller) patchData(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	id := mux.Vars(r)["id"]
	var input Employee

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		errorData := utils.ErrorLog{Skip: 1, Event: HttpReqReadErr, Message: err.Error()}
		utils.LogErrors(errorData)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, IncorrectInputErr)
		return
	}

	// need to have a single point of truth on id, should be the path var
	// or eliminate the path var
	input.ID = id

	err = c.validate.Struct(input)
	if err != nil {
		errorData := utils.ErrorLog{Skip: 1, Event: ValidationErr, Message: err.Error(), ErrorData: string(input.Json())}
		utils.LogErrors(errorData)

		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, IncorrectInputErr)
		return
	}

	updated, err := c.l.Update(input)
	if err != nil {
		if errors.Is(err, ErrDataConflict) {
			w.WriteHeader(http.StatusConflict)
			_, _ = fmt.Fprint(w, ErrDataConflict)
			return
		}
		if errors.Is(err, ErrDataNotFound) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = fmt.Fprint(w, ErrDataNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(updated)
	if err != nil {
		errorData := utils.ErrorLog{Skip: 1, Event: JsonEncodeErr, Message: err.Error(), ErrorData: string(updated.Json())}
		utils.LogErrors(errorData)
	}
}

func (c Controller) deleteData(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	id := mux.Vars(r)["id"]

	err := Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprint(w, ErrDataNotFound)
	}
}

func initService() {
	flagAppName := "Application name"
	flagLogLevel := "Application log-level"
	flagMock := "Enable data mocking"

	var appName string
	var logLevel string
	var mock bool

	flag.StringVar(&appName, "name", "apis", flagAppName)
	flag.StringVar(&appName, "n", "apis", flagAppName)
	flag.StringVar(&logLevel, "level", "info", flagLogLevel)
	flag.StringVar(&logLevel, "l", "info", flagLogLevel)
	flag.BoolVar(&mock, "mock", false, flagMock)
	flag.BoolVar(&mock, "m", false, flagMock)
	flag.Parse()

	ServiceInfo.NAME = appName
	ServiceInfo.ID = GetServiceId()

	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fields := Log.Fields{
		"hostname": hostName,
		"service":  ServiceInfo.NAME,
		"id":       ServiceInfo.ID,
	}

	var level Log.Level
	switch logLevel {
	case "debug":
		level = Log.DebugLevel
	case "error":
		level = Log.ErrorLevel
	case "fatal":
		level = Log.FatalLevel
	case "warn":
		level = Log.WarnLevel
	default:
		level = Log.InfoLevel
	}

	utils.InitLogs(fields, level)

	utils.Logger.WithFields(utils.StandardFields).WithFields(Log.Fields{"args": os.Args, "mode": "init", "logLevel": level}).Info("Service started successfully.")

	InitValidator()

	c = Controller{
		l:        &l,
		validate: Validate,
	}

	ic = InfoController{
		ServiceInfo: ServiceInfo,
	}

	l.serviceData = make(map[string]Employee)

	if mock {
		err := loadMockData()

		if err == nil {
			utils.Logger.Info("Mock data loaded successfully.")
		} else {
			errorData := utils.ErrorLog{Skip: 1, Event: MockDataErr, Message: err.Error()}
			utils.LogErrors(errorData)
		}
	}
}

func main() {
	initService()

	utils.Logger.WithFields(utils.StandardFields).WithFields(Log.Fields{"mode": "run"}).Info("Listening on port 8080")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/healthz", ic.healthCheck).Methods("GET")
	router.HandleFunc("/info", ic.getServiceInfo).Methods("GET")
	router.HandleFunc("/data", c.getAllData).Methods("GET")
	router.HandleFunc("/data", c.createData).Methods("PUT")
	router.HandleFunc("/data/{id}", c.getData).Methods("GET")
	router.HandleFunc("/data/{id}", c.patchData).Methods("PATCH")
	router.HandleFunc("/data/{id}", c.deleteData).Methods("DELETE")

	fmt.Println(http.ListenAndServe(":8080", router))
}
