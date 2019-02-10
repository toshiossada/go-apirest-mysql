package EmployeeRepository

import (
	"database/sql"
	"fmt"

	"github.com/toshiossada/go-restapi-mysql/models"
)

func dbConn() (db *sql.DB, err error) {
	config, _ := models.LoadConfiguration("config.json")

	connectionString := fmt.Sprint(config.Database.User, ":", config.Database.Password, "@tcp(", config.Database.Host, ")/", config.Database.DatabaseName)
	db, err = sql.Open(config.Database.Driver, connectionString)

	if err != nil {
		return nil, err
	}
	return db, nil
}

func ListAll() ([]models.Employee, error) {
	db, _ := dbConn()
	emp := models.Employee{}
	res := []models.Employee{}
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		return res, err
	}

	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			return res, err
		}
		emp.ID = id
		emp.Name = name
		emp.City = city
		res = append(res, emp)
	}

	defer db.Close()
	return res, nil
}

func GetById(nID int) (models.Employee, error) {
	db, _ := dbConn()
	emp := models.Employee{}

	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nID)
	if err != nil {
		return emp, err
	}

	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			return emp, err
		}
		emp.ID = id
		emp.Name = name
		emp.City = city
	}

	defer db.Close()

	return emp, nil
}

func Insert(emp models.Employee) (models.Employee, error) {
	db, _ := dbConn()

	insForm, err := db.Prepare("INSERT INTO Employee(name, city) VALUES(?,?)")
	if err != nil {
		return emp, err
	}
	res, err := insForm.Exec(emp.Name, emp.City)
	if err != nil {
		return emp, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return emp, err
	}
	emp.ID = int(id)

	defer db.Close()

	return emp, nil
}

func Update(emp models.Employee, nID int) (bool, error) {
	db, _ := dbConn()

	insForm, err := db.Prepare("UPDATE Employee SET name=?, city=? WHERE id=?")
	if err != nil {
		return false, err
	}
	_, err = insForm.Exec(emp.Name, emp.City, emp.ID)

	if err != nil {
		return false, err
	}

	defer db.Close()

	return true, nil
}

func Delete(nID int) (bool, error) {
	db, _ := dbConn()

	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")

	if err != nil {
		return false, err
	}
	_, err = delForm.Exec(nID)
	if err != nil {
		return false, err
	}

	defer db.Close()

	return true, nil
}

func Exists(id int) (bool, error) {
	db, _ := dbConn()
	selDB, err := db.Query("SELECT count(1) as count FROM Employee WHERE id=?", id)
	if err != nil {
		return false, err
	}

	count := 0

	for selDB.Next() {
		err = selDB.Scan(&count)
		println("count: ", count)
		if err != nil {
			return false, err
		}

	}

	defer db.Close()

	return count > 0, nil
}
