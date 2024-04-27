package client

type DummyIPC struct{}

func (d DummyIPC) Receive() ([]ReceivedData, error) {
	return []ReceivedData{}, nil
}

func (d DummyIPC) Dispatch(*ByteQueue) ([]byte, error) {
	return []byte{}, nil
}

func (d DummyIPC) Workspaces() ([]Workspace, error) {
	return []Workspace{}, nil
}

func (d DummyIPC) ActiveWorkspace() (Workspace, error) {
	return Workspace{}, nil
}

func (d DummyIPC) Monitors() ([]Monitor, error) {
	return []Monitor{}, nil
}

func (d DummyIPC) Clients() ([]Client, error) {
	return []Client{}, nil
}

func (d DummyIPC) ActiveWindow() (Window, error) {
	return Window{}, nil
}

func (d DummyIPC) Layers() (Layers, error) {
	return Layers{}, nil
}

func (d DummyIPC) Devices() (Devices, error) {
	return Devices{}, nil
}

func (d DummyIPC) Version() (Version, error) {
	return Version{}, nil
}

func (d DummyIPC) Keyword(*ByteQueue) error {
	return nil
}

func (d DummyIPC) Reload() error {
	return nil
}

func (d DummyIPC) SetCursor(string, string) error {
	return nil
}

func (d DummyIPC) Kill() error {
	return nil
}

func (d DummyIPC) Splash() (string, error) {
	return "", nil
}

func (d DummyIPC) GetOption(string) (string, error) {
	return "", nil
}

func (d DummyIPC) CursorPos() (CursorPos, error) {
	return CursorPos{}, nil
}

func (d DummyIPC) Binds() ([]Bind, error) {
	return []Bind{}, nil
}
