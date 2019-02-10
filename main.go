package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/toshiossada/go-restapi-mysql/models"
	"github.com/toshiossada/go-restapi-mysql/route"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, _ := models.LoadConfiguration("config.json")

	if numprocs := config.CPU.Maxprocs; numprocs != 0 {
		runtime.GOMAXPROCS(numprocs)
	}

	api := route.InitRouter()
	fmt.Println("Listening at ", config.Port, " port")
	http.ListenAndServe(config.Port, api.MakeHandler())
}
