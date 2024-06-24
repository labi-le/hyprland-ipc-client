# hypland ipc client

This is a client for the Hyprland IPC server. It is used to communicate with the Hyprland server.

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

### Notice
I stopped using Hyprland as my main wm and so I don't follow its development, hence the addition of which api functionality should not be expected, you can make a pull request or ask me to add something\
That doesn't mean I've abandoned the project, my attention is focused on other things, but I continue to support this project
