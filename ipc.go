package client

import (
	"encoding/json"
	"errors"
	"net"
	"strings"
)

const (
	BufSize   = 8192
	Separator = ">>"
)

type Client interface {
	Receive() ([]ReceivedData, error)
	Dispatch(args Args) ([]byte, error)
	Workspaces() ([]Workspace, error)
	// TODO: implement all (or most) of the ipc commands
	// https://wiki.hyprland.org/Configuring/Using-hyprctl/
	//Monitors() ([]Monitor, error)
	//Clients() ([]Client, error)
	//ActiveWindow() (Window, error)
	//Layers() ([]Layer, error)
	//Devices() ([]Device, error)
	//Keyword() (Keyword, error)
	//Version() (Version, error)
	//Kill() error
	//Splash() (string, error)
	//Reload() error
	//SetCursor(theme, size string) error
	//GetOption(name string) (string, error)
	//CursorPos() (CursorPos, error)
}

type Args struct {
	ArgvQueue
}

type ReceivedData struct {
	Type EventType
	Data string
}

type client struct {
	conn  net.Conn
	wconn *net.UnixAddr
	sign  string
}

func NewClient(sign string) Client {
	if sign == "" {
		panic("sign is empty")
	}

	conn, err := net.Dial("unix", "/tmp/hypr/"+sign+"/.socket2.sock")
	if err != nil {
		panic(err)
	}

	return &client{
		conn: conn,
		wconn: &net.UnixAddr{
			Net:  "unix",
			Name: "/tmp/hypr/" + sign + "/.socket.sock",
		},
		sign: sign,
	}
}

func (c *client) request(a Args) ([]byte, error) {
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

func (c *client) Receive() ([]ReceivedData, error) {
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

func (c *client) Dispatch(a Args) ([]byte, error) {
	if a.Len() == 0 {
		return nil, errors.New("wtfuq man you need to pass some args")
	}

	a.Add("dispatch")

	return c.request(a)
}

type Workspace struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Monitor         string `json:"monitor"`
	Windows         int    `json:"windows"`
	HasFullScreen   bool   `json:"hasfullscreen"`
	LastWindow      string `json:"lastwindow"`
	LastWindowTitle string `json:"lastwindowtitle"`
}

func (c *client) Workspaces() ([]Workspace, error) {
	a := Args{}
	a.Add("workspaces")

	buf, err := c.request(a)
	if err != nil {
		return nil, err
	}

	var ws []Workspace
	if err := json.Unmarshal(buf, &ws); err != nil {
		return nil, err
	}

	return ws, nil
}
