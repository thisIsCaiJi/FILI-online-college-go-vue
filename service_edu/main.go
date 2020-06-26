package main

import (
	"github.com/thisIsCaiJi/online_college/service_edu/handlers"
)

func main() {
	defer handlers.Close()
	handlers.ServerRun()
}
