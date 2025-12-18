package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func handleShutdown(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("shutdown", "/s", "/f", "/t", "0")
	err := cmd.Start()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprint(w, "Shutting down...")
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	// Safety check: Windows only
	if runtime.GOOS != "windows" {
		http.Error(w, "Logout only supported on Windows", 500)
		return
	}

	// shutdown /l = log out current user
	cmd := exec.Command("shutdown", "/l")
	err := cmd.Start()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, "Logging out...")
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
	http.HandleFunc("/logout", handleLogout)

	fmt.Println("Server running on :3000")
	http.ListenAndServe(":3000", nil)
}