package client

// EventHandler Hyprland will write to each connected ipc live events like this:
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
	Screencast(s Screencast)
}

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
	EventScreencast       EventType = "screencast"
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
		EventScreencast,
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

type Screencast struct {
	// True if a screen or window is being shared.
	Sharing bool

	// "0" if monitor is shared, "1" if window is shared.
	Owner string
}
