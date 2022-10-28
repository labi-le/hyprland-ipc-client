package client

type EventType string

const (
	EventWorkspace        EventType = "workspace"
	EventFocusedMonitor   EventType = "focusedmon"
	EventActiveWindow     EventType = "activewindow"
	EventFullscreen       EventType = "fullscreen"
	EventMonitorRemoved   EventType = "monitorremoved"
	EventMonitorAdded     EventType = "monitoradded"
	EventCreateWorkspace  EventType = "createworkspace"
	EventDestroyWorkspace EventType = "destroyworkspace"
	EventMoveWorkspace    EventType = "moveworkspace"
	EventActiveLayout     EventType = "activelayout"
	EventOpenWindow       EventType = "openwindow"
	EventCloseWindow      EventType = "closewindow"
	EventMoveWindow       EventType = "movewindow"
	EventOpenLayer        EventType = "openlayer"
	EventCloseLayer       EventType = "closelayer"
	EventSubMap           EventType = "submap"
)

func GetAllEvents() []EventType {
	return []EventType{
		EventWorkspace,
		EventFocusedMonitor,
		EventActiveWindow,
		EventFullscreen,
		EventMonitorRemoved,
		EventMonitorAdded,
		EventCreateWorkspace,
		EventDestroyWorkspace,
		EventMoveWorkspace,
		EventActiveLayout,
		EventOpenWindow,
		EventCloseWindow,
		EventMoveWindow,
		EventOpenLayer,
		EventCloseLayer,
		EventSubMap,
	}
}

type MoveWorkspace struct {
	WorkspaceName
	MonitorName
}

type MonitorName string

type FocusedMonitor struct {
	MonitorName
	WorkspaceName
}

type WorkspaceName string

type SubMap string

type CloseLayer string

type OpenLayer string

type MoveWindow struct {
	Address string
	WorkspaceName
}

type CloseWindow struct {
	Address string
}

type OpenWindow struct {
	Address, Class, Title string
	WorkspaceName
}

type ActiveLayout struct {
	Type, Name string
}

type ActiveWindow struct {
	Name, Title string
}

type ActiveWorkspace WorkspaceName
