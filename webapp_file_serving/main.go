package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os/exec"
	"runtime"
)

//go:embed static
var staticFiles embed.FS

type Post struct {
	Title   string
	Content string
}

// opens url in web browser
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func main() {
	// http.HandleFunc("/", indexHandler)
	staticFS := fs.FS(staticFiles)
	htmlContent, _ := fs.Sub(staticFS, "static")
	fileSys := http.FileServer(http.FS(htmlContent))
	http.Handle("/", fileSys)

	go open("http://localhost:3000")
	panic(http.ListenAndServe(":3000", nil))
}
