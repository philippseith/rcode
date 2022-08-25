package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprint(flag.CommandLine.Output(),
			`rcode is a tool for working with remotes and Visual Studio Code. 
It helps you if you are connected via terminal and you want to edit the current folder with vscode on your host.

Usage:
Run 'rcode -address host:port -remote someremote' on the the host and
'curl -d "path" -X PUT http://host:port' on the remote named 'someremote' in your vscode Remote Explorer, 
where 'path' is the path you want to open in vscode.
`)
		flag.PrintDefaults()
	}
	var address, remoteName, code string
	flag.StringVar(&address, "address", "10.0.0.2:49374", "address to bind the server to")
	flag.StringVar(&remoteName, "remoteName", "", "name of the remote in vscode remote explorer")
	flag.StringVar(&code, "code", "code", "how to invoke Visual Studio Code")
	flag.Parse()

	http.HandleFunc("/api/rcode", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPut {
			body, err := io.ReadAll(request.Body)
			if err != nil {
				request.Response.StatusCode = http.StatusInternalServerError
			}
			_ = request.Body.Close()
			remotePath := string(body)
			exec.Command(code, "--folder-uri",
				fmt.Sprintf("vscode://ssh-remote%%2B%s/%s", remoteName, remotePath))
		}
	})
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err)
	}
}
