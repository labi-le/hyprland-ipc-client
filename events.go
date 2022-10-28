package client

import (
	"strings"
)

// EventHandler Hyprland will write to each connected client live events like this:
type EventHandler interface {
	// Workspace emitted on workspace change. Is emitted ONLY when a user requests a workspace change, and is not emitted on mouse movements (see activemon)
	Workspace(w WorkspaceName)
	// FocusedMonitor emitted on the active monitor being changed.
	FocusedMonitor(m FocusedMonitor)
	// ActiveWindow emitted on the active window being changed.
	ActiveWindow(w ActiveWindow)
	// Fullscreen emitted when a fullscreen status of a window changes.
	Fullscreen(f bool)
	// MonitorRemoved emitted when a monitor is removed (disconnected)
	MonitorRemoved(m MonitorName)
	// MonitorAdded emitted when a monitor is added (connected)
	MonitorAdded(m MonitorName)
	// CreateWorkspace emitted when a workspace is created
	CreateWorkspace(w WorkspaceName)
	// DestroyWorkspace emitted when a workspace is destroyed
	DestroyWorkspace(w WorkspaceName)
	// MoveWorkspace emitted when a workspace is moved to a different monitor
	MoveWorkspace(w MoveWorkspace)
	// ActiveLayout emitted on a layout change of the active keyboard
	ActiveLayout(l ActiveLayout)
	// OpenWindow emitted when a window is opened
	OpenWindow(o OpenWindow)
	// CloseWindow emitted when a window is closed
	CloseWindow(c CloseWindow)
	// MoveWindow emitted when a window is moved to a workspace
	MoveWindow(m MoveWindow)
	// OpenLayer emitted when a layerSurface is mapped
	OpenLayer(l OpenLayer)
	// CloseLayer emitted when a layerSurface is unmapped
	CloseLayer(c CloseLayer)
	// SubMap emitted when a keybind submap changes. Empty means default.
	SubMap(s SubMap)
}

func Subscribe(c Client, ev EventHandler, events ...EventType) error {
	for {
		msg, err := c.Receive()
		if err != nil {
			return err
		}

		for _, data := range msg {
			processEvent(ev, data, events)
		}
	}
}

func processEvent(ev EventHandler, msg ReceivedData, events []EventType) {
	for _, event := range events {
		raw := strings.Split(msg.Data, ",")
		if msg.Type == event {
			switch event {
			case EventWorkspace:
				// e.g. "1" (workspace number)
				ev.Workspace(WorkspaceName(raw[0]))
				break
			case EventFocusedMonitor:
				// idk
				ev.FocusedMonitor(FocusedMonitor{
					MonitorName:   MonitorName(raw[0]),
					WorkspaceName: WorkspaceName(raw[1]),
				})
				break
			case EventActiveWindow:
				// e.g. jetbrains-goland,hyprland-ipc-client – main.go
				ev.ActiveWindow(ActiveWindow{
					Name:  raw[0],
					Title: raw[1],
				})

				break
			case EventFullscreen:
				// e.g. "true" or "false"
				ev.Fullscreen(raw[0] == "1")
				break
			case EventMonitorRemoved:
				// e.g. idk
				ev.MonitorRemoved(MonitorName(raw[0]))
				break
			case EventMonitorAdded:
				// e.g. idk
				ev.MonitorAdded(MonitorName(raw[0]))
				break
			case EventCreateWorkspace:
				// e.g. "1" (workspace number)
				ev.CreateWorkspace(WorkspaceName(raw[0]))
				break
			case EventDestroyWorkspace:
				// e.g. "1" (workspace number)
				ev.DestroyWorkspace(WorkspaceName(raw[0]))
				break
			case EventMoveWorkspace:
				// e.g. idk
				ev.MoveWorkspace(MoveWorkspace{
					WorkspaceName: WorkspaceName(raw[0]),
					MonitorName:   MonitorName(raw[1]),
				})
				break
			case EventActiveLayout:
				// e.g. AT Translated Set 2 keyboard,Russian
				ev.ActiveLayout(ActiveLayout{
					Type: raw[0],
					Name: raw[1],
				})
				break
			case EventOpenWindow:
				// e.g. 80864f60,1,Alacritty,Alacritty
				ev.OpenWindow(OpenWindow{
					Address:       raw[0],
					WorkspaceName: WorkspaceName(raw[1]),
					Class:         raw[2],
					Title:         raw[3],
				})
				break
			case EventCloseWindow:
				// e.g. 5
				ev.CloseWindow(CloseWindow{
					Address: raw[0],
				})
				break
			case EventMoveWindow:
				// e.g. 5
				ev.MoveWindow(MoveWindow{
					Address:       raw[0],
					WorkspaceName: WorkspaceName(raw[1]),
				})
				break
			case EventOpenLayer:
				// e.g. wofi
				ev.OpenLayer(OpenLayer(raw[0]))
				break
			case EventCloseLayer:
				// e.g. wofi
				ev.CloseLayer(CloseLayer(raw[0]))
				break
			case EventSubMap:
				// e.g. idk
				ev.SubMap(SubMap(raw[0]))
				break
			}

		}
	}
}

func MakeDummyEvHandler() EventHandler {
	return &DummyEvHandler{}
}

type DummyEvHandler struct{}

func (e *DummyEvHandler) Workspace(WorkspaceName)        {}
func (e *DummyEvHandler) FocusedMonitor(FocusedMonitor)  {}
func (e *DummyEvHandler) ActiveWindow(ActiveWindow)      {}
func (e *DummyEvHandler) Fullscreen(bool)                {}
func (e *DummyEvHandler) MonitorRemoved(MonitorName)     {}
func (e *DummyEvHandler) MonitorAdded(MonitorName)       {}
func (e *DummyEvHandler) CreateWorkspace(WorkspaceName)  {}
func (e *DummyEvHandler) DestroyWorkspace(WorkspaceName) {}
func (e *DummyEvHandler) MoveWorkspace(MoveWorkspace)    {}
func (e *DummyEvHandler) ActiveLayout(ActiveLayout)      {}
func (e *DummyEvHandler) OpenWindow(OpenWindow)          {}
func (e *DummyEvHandler) CloseWindow(CloseWindow)        {}
func (e *DummyEvHandler) MoveWindow(MoveWindow)          {}
func (e *DummyEvHandler) OpenLayer(OpenLayer)            {}
func (e *DummyEvHandler) CloseLayer(CloseLayer)          {}
func (e *DummyEvHandler) SubMap(SubMap)                  {}
