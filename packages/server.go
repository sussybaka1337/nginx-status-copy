package packages

import (
	"fmt"
	"net/http"
	"server_status/routes"
)

func StartServer() {
	for pattern, handler := range routes.GetAllRoutes() {
		http.HandleFunc(pattern, handler)
	}
	if err := http.ListenAndServe(":8999", nil); err != nil {
		fmt.Println(err)
	}
}
