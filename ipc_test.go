package client

import (
	"os"
	"reflect"
	"testing"
)

var ipctest = NewClient(os.Getenv("HYPRLAND_INSTANCE_SIGNATURE"))

func Test_ipc_Clients(t *testing.T) {
	got, err := ipctest.Clients()
	if err != nil {
		t.Error(err)
	}

	if len(got) == 0 {
		t.Error("got is empty")
	}

	for _, client := range got {
		if reflect.DeepEqual(client, Client{}) {
			t.Errorf("got empty struct")
		}
	}
}

func Test_ipc_Workspaces(t *testing.T) {
	got, err := ipctest.Workspaces()
	if err != nil {
		t.Error(err)
	}

	if len(got) == 0 {
		t.Error("got is empty")
	}

	for _, workspace := range got {
		if reflect.DeepEqual(workspace, Workspace{}) {
			t.Errorf("got empty struct")
		}
	}
}

func Test_ipc_ActiveWorkspaces(t *testing.T) {
	got, err := ipctest.ActiveWorkspace()
	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(got, Workspace{}) {
		t.Errorf("got empty struct")
	}
}

func Test_ipc_Monitors(t *testing.T) {
	got, err := ipctest.Monitors()

	if err != nil {
		t.Error(err)
	}

	if len(got) == 0 {
		t.Error("got is empty")
	}

	for _, workspace := range got {
		if reflect.DeepEqual(workspace, Workspace{}) {
			t.Errorf("got empty struct")
		}
	}
}

func Test_ipc_ActiveWindow(t *testing.T) {
	got, err := ipctest.ActiveWindow()
	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(got, ActiveWindow{}) {
		t.Errorf("got empty struct")
	}
}

func Test_ipc_Layer(t *testing.T) {
	got, err := ipctest.Layers()
	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(got, Layer{}) {
		t.Errorf("got empty struct")
	}
}

func Test_ipc_Devices(t *testing.T) {
	got, err := ipctest.Devices()
	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(got, Devices{}) {
		t.Errorf("got empty struct")
	}
}

func Test_ipc_Version(t *testing.T) {
	got, err := ipctest.Version()
	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(got, Version{}) {
		t.Errorf("got empty struct")
	}
}

func Test_ipc_Keyword(t *testing.T) {
	q := NewByteQueue()
	q.Add([]byte("general:border_size 1"))

	err := ipctest.Keyword(q)
	if err != nil {
		t.Error(err)
	}
}

func Test_ipc_Kill(t *testing.T) {
	//err := ipctest.Kill()
	//if err != nil {
	//	t.Error(err)
	//}
	//
}

func Test_ipc_Reload(t *testing.T) {
	err := ipctest.Reload()
	if err != nil {
		t.Error(err)
	}
}

func Test_ipc_Splash(t *testing.T) {
	got, err := ipctest.Splash()
	if err != nil {
		t.Error(err)
	}

	if got == "" {
		t.Error("got is empty")
	}
}

func Test_ipc_SetCursor(t *testing.T) {
	err := ipctest.SetCursor("Bibata-Modern-Classic", "24")
	if err != nil {
		t.Error(err)
	}
}

func Test_ipc_GetOption(t *testing.T) {
	got, err := ipctest.GetOption("general:border_size")
	if err != nil {
		t.Error(err)
	}

	if got == "" {
		t.Error("got is empty")
	}
}

func Test_ipc_CursorPos(t *testing.T) {
	got, err := ipctest.CursorPos()
	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(got, CursorPos{}) {
		t.Errorf("got empty struct")
	}
}

func TestIpc_Dispatch(t *testing.T) {
	q := NewByteQueue()
	q.Add([]byte("exec"))
	_, err := ipctest.Dispatch(q)
	if err != nil {
		t.Error(err)
	}
}
