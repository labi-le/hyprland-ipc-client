# hypland ipc client

This is a client for the Hypland IPC server. It is used to communicate with the Hypland server.

## Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/labi-le/hyprland-ipc-client"
)

type ed struct {
	client.DummyEvHandler
}

func main() {
	c := client.NewClient(os.Getenv("HYPRLAND_INSTANCE_SIGNATURE"))
	e := client.MakeDummyEvHandler()
	client.Subscribe(c, e, client.EventActiveLayout)
}

func (e *ed) ActiveLayout(layout client.ActiveLayout) {
	fmt.Println("ActiveLayout", layout.Type, layout.Name)
}

```