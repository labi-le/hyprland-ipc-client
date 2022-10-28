package client

import "testing"

var (
	h = &FakeEventHandler{}
	c = &FakeClient{}
)

type FakeClient struct {
}

type FakeEventHandler struct {
	DummyEvHandler
}

func (f FakeClient) Receive() ([]ReceivedData, error) {
	return []ReceivedData{
		{
			Type: EventWorkspace,
			Data: "1",
		},

		{
			Type: EventFocusedMonitor,
			Data: "1,1",
			// TODO I only have one monitor, so I didn't check this
		},

		{
			Type: EventActiveWindow,
			Data: "jetbrains-goland,hyprland-ipc-client – ipc.go",
		},

		{
			Type: EventFullscreen,
			Data: "1",
		},

		{
			Type: EventMonitorRemoved,
			Data: "1",
			// TODO I only have one monitor, so I didn't check this
		},

		{
			Type: EventMonitorAdded,
			Data: "1",
			// TODO I only have one monitor, so I didn't check this
		},

		{
			Type: EventCreateWorkspace,
			Data: "1",
		},

		{
			Type: EventDestroyWorkspace,
			Data: "1",
		},

		{
			Type: EventMoveWorkspace,
			Data: "1,1",
			// TODO I only have one monitor, so I didn't check this
		},

		{
			Type: EventActiveLayout,
			Data: "AT Translated Set 2 keyboard,Russian",
		},

		{
			Type: EventOpenWindow,
			Data: "80e62df0,2,jetbrains-goland,win430",
		},

		{
			Type: EventCloseWindow,
			Data: "80e62df0",
		},

		{
			Type: EventMoveWindow,
			Data: "80e62df0,1",
		},

		{
			Type: EventOpenLayer,
			Data: "wofi",
		},

		{
			Type: EventCloseLayer,
			Data: "wofi",
		},

		{
			Type: EventSubMap,
			Data: "1",
			// idk
		},
	}, nil
}

func (f FakeClient) Dispatch(Args) ([]byte, error) {
	panic("needlessly")
}

func (f FakeClient) Workspaces() ([]Workspace, error) {
	panic("needlessly")
}

func TestSubscribe(t *testing.T) {
	err := SubscribeWithoutLoop(c, h, GetAllEvents()...)
	if err != nil {
		t.Error(err)
	}
}

func SubscribeWithoutLoop(c Client, ev EventHandler, events ...EventType) error {
	msg, err := c.Receive()
	if err != nil {
		return err
	}

	for _, data := range msg {
		processEvent(ev, data, events)
	}

	return nil
}
