# hypland ipc client

This is a client for the Hypland IPC server. It is used to communicate with the Hypland server.

### Installation
```sh
go get -u github.com/labi-le/hyprland-ipc-client/v3
```

## Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/labi-le/hyprland-ipc-client/v3"
)

type ed struct {
	client.DummyEvHandler
}

func main() {
	c := client.MustClient(os.Getenv("HYPRLAND_INSTANCE_SIGNATURE"))
	e := &ed{}
	client.Subscribe(c, e, client.EventActiveLayout)
}

func (e *ed) ActiveLayout(layout client.ActiveLayout) {
	fmt.Println("ActiveLayout", layout.Type, layout.Name)
}

```
