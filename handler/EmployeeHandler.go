package EmployeeHandler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/toshiossada/go-restapi-mysql/models"
	EmployeeRepository "github.com/toshiossada/go-restapi-mysql/repository"
)

func Serialize(w rest.ResponseWriter, r *rest.Request) (models.Employee, error) {
	employee := models.Employee{}
	if err := r.DecodeJsonPayload(&employee); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return employee, err
	}

	return employee, nil
}

func ListAll(w rest.ResponseWriter, r *rest.Request) {

	res, err := EmployeeRepository.ListAll()

	if err != nil {
		log.Println(err.Error())
		rest.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.WriteJson(&res)
}

func GetById(w rest.ResponseWriter, r *rest.Request) {
	nID, _ := strconv.Atoi(r.PathParam("id"))

	exist, _ := EmployeeRepository.Exists(nID)
	fmt.Println(exist)
	if !exist {
		log.Println(http.StatusText(http.StatusNotFound))
		rest.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	emp, err := EmployeeRepository.GetById(nID)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.WriteJson(&emp)
}

func Insert(w rest.ResponseWriter, r *rest.Request) {

	emp, err := Serialize(w, r)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	emp, err = EmployeeRepository.Insert(emp)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(fmt.Sprint("INSERT: Name: ", emp.Name, " | City: ", emp.City))

	w.WriteHeader(http.StatusCreated)
	w.WriteJson(&emp)
}

func Update(w rest.ResponseWriter, r *rest.Request) {

	nID, _ := strconv.Atoi(r.PathParam("id"))
	emp, err := Serialize(w, r)
	log.Println(fmt.Sprint("UPDATE: Name: ", emp.Name, " | City: ", emp.City))

	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if nID != emp.ID {
		log.Println(http.StatusText(http.StatusBadRequest))
		rest.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	exist, _ := EmployeeRepository.Exists(nID)
	fmt.Println(exist)
	if !exist {
		log.Println(http.StatusText(http.StatusNotFound))
		rest.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	result, err := EmployeeRepository.Update(emp, nID)

	if err != nil {
		log.Println(err.Error())
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !result {
		log.Println(http.StatusText(http.StatusBadRequest))
		rest.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func Delete(w rest.ResponseWriter, r *rest.Request) {
	nID, _ := strconv.Atoi(r.PathParam("id"))
	log.Println(fmt.Sprint("DELETE ID: ", nID))

	exist, _ := EmployeeRepository.Exists(nID)
	fmt.Println(exist)
	if !exist {
		log.Println(http.StatusText(http.StatusNotFound))
		rest.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	result, err := EmployeeRepository.Delete(nID)
	if err != nil {
		log.Println(err.Error())
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !result {
		log.Println(http.StatusText(http.StatusBadRequest))
		rest.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
