package client

import (
	"encoding/json"
	"errors"
	"io"
	"net"
	"reflect"
	"strings"
)

const (
	BufSize   = 8192
	Separator = ">>"
	Success   = "ok"
)

type ReceivedData struct {
	Type EventType
	Data string
}

type ipc struct {
	conn  net.Conn
	wconn *net.UnixAddr
	sign  string
}

func NewClient(sign string) IPC {
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

func (c *ipc) request(q *ByteQueue) ([]byte, error) {
	if q.Len() == 0 {
		return nil, errors.New("wtfuq man you need to pass some args")
	}

	conn, err := net.DialUnix("unix", nil, c.wconn)
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	if q.Len() > 1 {
		q.Back([]byte("[[BATCH]]"))
	}

	glued := q.Glue()

	req := Get(len(glued) + 2)
	copy(req, "j/")
	copy(req[2:], glued)

	defer Put(req)

	if _, err := conn.Write(req); err != nil {
		return nil, err
	}

	var response []byte
	buf := Get(BufSize)
	defer Put(buf)

	for {
		n, tcpErr := conn.Read(buf)
		if tcpErr != nil {
			if tcpErr == io.EOF {
				break
			}
			return nil, tcpErr
		}

		response = append(response, buf[:n]...)

		if n < BufSize {
			break
		}
	}

	return response, nil
}

// wrapreq
// a command without arguments can be safely wrapped in one method so as not to write the same thing every time
// v is a pointer to a struct
func (c *ipc) wrapreq(command string, v any, q *ByteQueue) error {
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		panic("v must be a pointer to a structure")
	}

	if q == nil {
		q = NewByteQueue()
	}

	q.Add(UnsafeBytes(command))

	buf, err := c.request(q)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(buf, v); err != nil {
		return err
	}

	return nil
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

func (c *ipc) Dispatch(a *ByteQueue) ([]byte, error) {
	a.Add([]byte("dispatch"))

	return c.request(a)
}

func (c *ipc) Workspaces() ([]Workspace, error) {
	var workspaces []Workspace
	return workspaces, c.wrapreq("workspaces", &workspaces, nil)
}

func (c *ipc) ActiveWorkspace() (Workspace, error) {
	var workspace Workspace
	return workspace, c.wrapreq("activeworkspace", &workspace, nil)
}

func (c *ipc) Monitors() ([]Monitor, error) {
	var monitors []Monitor
	return monitors, c.wrapreq("monitors", &monitors, nil)
}

func (c *ipc) Clients() ([]Client, error) {
	var clients []Client
	return clients, c.wrapreq("clients", &clients, nil)
}

func (c *ipc) ActiveWindow() (Window, error) {
	var window Window
	return window, c.wrapreq("activewindow", &window, nil)
}

func (c *ipc) Layers() (Layers, error) {
	var layers Layers
	return layers, c.wrapreq("layers", &layers, nil)
}

func (c *ipc) Devices() (Devices, error) {
	var devices Devices
	return devices, c.wrapreq("devices", &devices, nil)
}

func (c *ipc) Keyword(args *ByteQueue) error {
	args.Back([]byte("keyword"))

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
	var version Version
	return version, c.wrapreq("version", &version, nil)
}

func (c *ipc) Kill() error {
	q := NewByteQueue()
	q.Add([]byte("kill"))

	_, err := c.Dispatch(q)
	return err
}

func (c *ipc) Reload() error {
	q := NewByteQueue()
	q.Add([]byte("reload"))

	_, err := c.request(q)
	return err
}

func (c *ipc) SetCursor(theme, size string) error {
	q := NewByteQueue()
	q.Add(UnsafeBytes(theme))
	q.Add(UnsafeBytes(size))
	q.Back([]byte("setcursor"))

	_, err := c.request(q)
	return err
}

func (c *ipc) GetOption(name string) (string, error) {
	q := NewByteQueue()
	q.Add(UnsafeBytes(name))
	q.Back([]byte("getoption"))

	buf, err := c.request(q)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func (c *ipc) Splash() (string, error) {
	q := NewByteQueue()
	q.Back([]byte("splash"))

	buf, err := c.request(q)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func (c *ipc) CursorPos() (CursorPos, error) {
	var cursorpos CursorPos
	return cursorpos, c.wrapreq("cursorpos", &cursorpos, nil)
}
