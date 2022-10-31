package client

import (
	"encoding/json"
	"errors"
	"net"
	"reflect"
	"strings"
)

const (
	BufSize   = 8192
	Separator = ">>"
	Success   = "ok"
)

type Args struct {
	ArgvQueue
}

type ReceivedData struct {
	Type EventType
	Data string
}

type ipc struct {
	conn  net.Conn
	wconn *net.UnixAddr
	sign  string
}

func NewIPC(sign string) IPC {
	if sign == "" {
		panic("sign is empty")
	}

	conn, err := net.Dial("unix", "/tmp/hypr/"+sign+"/.socket2.sock")
	if err != nil {
		panic(err)
	}

	return &ipc{
		conn: conn,
		wconn: &net.UnixAddr{
			Net:  "unix",
			Name: "/tmp/hypr/" + sign + "/.socket.sock",
		},
		sign: sign,
	}
}

func (c *ipc) request(a Args) ([]byte, error) {
	if a.Len() == 0 {
		return nil, errors.New("wtfuq man you need to pass some args")
	}

	conn, err := net.DialUnix("unix", nil, c.wconn)
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	var argv string
	if a.Len() > 1 {
		argv = "[[BATCH]] "
	}

	argv += "j/" + a.String()

	if _, err := conn.Write([]byte(argv)); err != nil {
		return nil, err
	}

	var buf = make([]byte, BufSize)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf[:n], nil
}

// wrapreq
// a command without arguments can be safely wrapped in one method so as not to write the same thing every time
// v is a pointer to a struct
func (c *ipc) wrapreq(command string, v any, a Args) (any, error) {
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		panic("v must be a pointer to a structure")
	}

	a.Push(command)

	buf, err := c.request(a)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(buf, v); err != nil {
		return nil, err
	}

	// cast to the desired type
	return reflect.ValueOf(v).Elem().Interface(), nil
}

func (c *ipc) Receive() ([]ReceivedData, error) {
	buf := make([]byte, BufSize)
	n, err := c.conn.Read(buf)
	if err != nil {
		return nil, err
	}

	buf = buf[:n]

	var recv []ReceivedData
	//EVENT>>DATA\n (\n is a linebreak)
	rawEvents := strings.Split(string(buf), "\n")
	for _, event := range rawEvents {
		if event == "" {
			continue
		}

		split := strings.Split(event, Separator)
		if split[0] == "" || split[1] == "" || split[1] == "," {
			continue
		}

		recv = append(recv, ReceivedData{
			Type: EventType(split[0]),
			Data: split[1],
		})
	}

	return recv, nil
}

func (c *ipc) Dispatch(a Args) ([]byte, error) {
	a.Push("dispatch")

	return c.request(a)
}

func (c *ipc) Workspaces() ([]Workspace, error) {
	wrapreq, err := c.wrapreq("workspaces", &[]Workspace{}, Args{})
	return wrapreq.([]Workspace), err
}

func (c *ipc) Monitors() ([]Monitor, error) {
	wrapreq, err := c.wrapreq("monitors", &[]Monitor{}, Args{})
	return wrapreq.([]Monitor), err
}

func (c *ipc) Clients() ([]Client, error) {
	wrapreq, err := c.wrapreq("clients", &[]Client{}, Args{})
	return wrapreq.([]Client), err
}

func (c *ipc) ActiveWindow() (Window, error) {
	wrapreq, err := c.wrapreq("activewindow", &Window{}, Args{})
	return wrapreq.(Window), err
}

func (c *ipc) Layers() (Layers, error) {
	wrapreq, err := c.wrapreq("layers", &Layers{}, Args{})
	return wrapreq.(Layers), err
}

func (c *ipc) Devices() (Devices, error) {
	wrapreq, err := c.wrapreq("devices", &Devices{}, Args{})
	return wrapreq.(Devices), err
}

func (c *ipc) Keyword(args Args) error {
	args.PushBack("keyword")

	response, err := c.request(args)
	if err != nil {
		return err
	}

	if ok := string(response); ok != Success {
		return errors.New(ok)
	}

	return nil
}

func (c *ipc) Version() (Version, error) {
	wrapreq, err := c.wrapreq("version", &Version{}, Args{})
	return wrapreq.(Version), err
}

func (c *ipc) Kill() error {
	a := Args{}
	a.Push("kill")

	_, err := c.Dispatch(a)
	return err
}

func (c *ipc) Reload() error {
	a := Args{}
	a.Push("reload")

	_, err := c.request(a)
	return err
}

func (c *ipc) SetCursor(theme, size string) error {
	a := Args{}
	a.Push(theme)
	a.Push(size)
	a.PushBack("setcursor")

	_, err := c.request(a)
	return err
}

func (c *ipc) GetOption(name string) (string, error) {
	a := Args{}
	a.Push(name)
	a.PushBack("getoption")

	buf, err := c.request(a)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func (c *ipc) Splash() (string, error) {
	a := Args{}
	a.PushBack("splash")

	buf, err := c.request(a)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func (c *ipc) CursorPos() (CursorPos, error) {
	wrapreq, err := c.wrapreq("cursorpos", &CursorPos{}, Args{})
	return wrapreq.(CursorPos), err
}
