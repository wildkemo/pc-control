package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func handleShutdown(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("shutdown", "/s", "/t", "0")
	err := cmd.Start()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprint(w, "Shutting down...")
}

func handleCloseAll(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("taskkill", "/F", "/IM", "notepad.exe") // change as needed
	err := cmd.Run()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprint(w, "Closed all notepad windows")
}

func main() {
	// get directory of the executable
	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(execPath)

	// serve static files from dist subdirectory
	distDir := filepath.Join(dir, "dist")
	fs := http.FileServer(http.Dir(distDir))
	http.Handle("/", fs)

	// endpoints for commands
	http.HandleFunc("/shutdown", handleShutdown)
	http.HandleFunc("/close-all", handleCloseAll)

	fmt.Println("Server running on :3000")
	http.ListenAndServe(":3000", nil)
}