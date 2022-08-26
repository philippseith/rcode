# rcode

`rcode` is a tool for working with remotes and Visual Studio Code. 

## Use Case 
You are connected via terminal and you want to edit the current folder with vscode on your host.

## Installation
`go install github.com/philippseith/rcode/cmd/rcode@latest`

## Usage
Run `rcode -address host:port -remote someremote` on the the host and
just `curl -d "path" -X PUT http://host:port/api/rcode` on the remote named `someremote` in your vscode Remote Explorer,
where `path` is the absolute path you want to open in vscode.
A new vscod window with the remote folder will be opened.

## Next
It should be fairly simple to write some shell scripts/functions to simplify sending the `PUT` with the path
