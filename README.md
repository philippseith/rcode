# rcode

`rcode` is a tool for working with remotes and Visual Studio Code. 

## Use Case 
You are connected via terminal and you want to edit the current folder with vscode on your host.

## Prerequisites

On the host:
- go 1.13 or higher

On the client:
- curl

## Installation

On the host:
- `go install github.com/philippseith/rcode/cmd/rcode@latest`

On the client:
- modify `rcode.sh` for your needs and put it into your path
## Usage
Run `rcode -address host:port -remote someremote` on the host and
change to the desired path on the client 
(the machine which is named `someremote` in the vscode Remote Explorer on the host) and run `rcode .`

A new vscode window with the remote folder will be opened.
