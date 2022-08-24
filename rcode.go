package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
)

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(),
			`%s is a small web server which listens for a PUT on/api/code/
and expects a folder on a remote. 
It will then try to open this folder on the remote.

Usage:
`, path.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	var address, remoteName string
	flag.StringVar(&address, "address", "10.0.0.2:49374", "address to bind the server to")
	flag.StringVar(&remoteName, "remoteName", "", "name of the remote in vscode remote explorer")

	flag.Parse()

	http.HandleFunc("/api/rcode", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPut {
			body, err := io.ReadAll(request.Body)
			if err != nil {
				request.Response.StatusCode = http.StatusInternalServerError
			}
			_ = request.Body.Close()
			remotePath := string(body)
			exec.Command("code", "--folder-uri",
				fmt.Sprintf("vscode://ssh-remote%%2B%s/%s", remoteName, remotePath))
		}
	})
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err)
	}
}
