package routes

import (
	"net/http"
	"sync"
)

func GetAllRoutes() map[string]func(http.ResponseWriter, *http.Request) {
	locker := new(sync.Mutex)
	go PollingRequestsPerSecond(locker)
	return map[string]func(http.ResponseWriter, *http.Request){
		"/": displayStatus,
	}
}
