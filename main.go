package main

import (
	"runtime"
	"server_status/packages"
)

func main() {
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)
	packages.StartServer()
}
