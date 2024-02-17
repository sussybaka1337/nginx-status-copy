package routes

import (
	"encoding/json"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

var requests uint64
var requestsPerSecond uint64

type ServerStatus struct {
	TotalRequests     uint64 `json:"total_requests"`
	RequestsPerSecond uint64 `json:"requests_per_second"`
}

func PollingRequestsPerSecond(locker *sync.Mutex) {
	for {
		currentRequests := requests
		time.Sleep(1 * time.Second)
		nextRequests := requests
		locker.Lock()
		requestsPerSecond = nextRequests - currentRequests
		locker.Unlock()
	}
}

func displayStatus(response http.ResponseWriter, request *http.Request) {
	atomic.AddUint64(&requests, 1)
	response.Header().Set("Content-Type", "application/json")
	status := ServerStatus{
		TotalRequests:     requests,
		RequestsPerSecond: requestsPerSecond,
	}
	json.NewEncoder(response).Encode(status)
}
