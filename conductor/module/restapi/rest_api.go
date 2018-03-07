package restapi

import (
	"log"
	"net/http"
	"github.com/emicklei/go-restful"
)

func Start() {
	user := UserResource{}
	restful.DefaultContainer.Add(user.WebService())

	log.Printf("start listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
