package main

import (
	"net/http"
	"os/exec"
	"syscall"

	"bey.com/routes"
)

func openBrowser() {
	cmd := exec.Command(`cmd`, `/c`, `start`, `http://localhost:3000`)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()
}

func main() {
	http.HandleFunc("/", routes.HomeHandler)
	http.HandleFunc("/list", routes.ListHandler)
	go openBrowser()
	http.ListenAndServe(":3000", nil)
}
