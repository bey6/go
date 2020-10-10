package main

import (
	"net/http"
	"os/exec"
	"syscall"
)

func openBrowser() {
	cmd := exec.Command(`cmd`, `/c`, `start`, `http://localhost:3000`)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/list", listHandler)
	go openBrowser()
	http.ListenAndServe(":3000", nil)
}
