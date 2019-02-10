package route

import (
	"log"

	"github.com/ant0ine/go-json-rest/rest"
	EmployeeHandler "github.com/toshiossada/go-restapi-mysql/handler"
)

func InitRouter() *rest.Api {

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/persons", EmployeeHandler.ListAll),
		rest.Post("/persons", EmployeeHandler.Insert),
		rest.Put("/persons/:id", EmployeeHandler.Update),
		rest.Get("/persons/:id", EmployeeHandler.GetById),
		rest.Delete("/persons/:id", EmployeeHandler.Delete),
	)

	if err != nil {
		log.Println("Erro")
	}

	api.SetApp(router)

	return api
}
