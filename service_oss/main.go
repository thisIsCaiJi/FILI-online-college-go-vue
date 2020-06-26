package main

import (
	"github.com/thisIsCaiJi/online_college/service_oss/handlers"
)

func main() {
	defer handlers.Close()
	handlers.ServerRun()
}
