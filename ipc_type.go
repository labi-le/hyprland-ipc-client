package client

type IPC interface {
	Receive() ([]ReceivedData, error)
	Dispatch(args *ByteQueue) ([]byte, error)
	Workspaces() ([]Workspace, error)
	ActiveWorkspace() (Workspace, error)
	Monitors() ([]Monitor, error)
	Clients() ([]Client, error)
	ActiveWindow() (Window, error)
	Layers() (Layers, error)
	Devices() (Devices, error)
	Keyword(args *ByteQueue) error
	Version() (Version, error)
	Kill() error
	Splash() (string, error)
	Reload() error
	SetCursor(theme, size string) error
	GetOption(name string) (string, error)
	CursorPos() (CursorPos, error)
	Binds() ([]Bind, error)
}

type Workspace struct {
	WorkspaceType
	Monitor         string `json:"monitor"`
	Windows         int    `json:"windows"`
	HasFullScreen   bool   `json:"hasfullscreen"`
	LastWindow      string `json:"lastwindow"`
	LastWindowTitle string `json:"lastwindowtitle"`
}

type WorkspaceType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Monitor struct {
	Id              int           `json:"id"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	Width           int           `json:"width"`
	Height          int           `json:"height"`
	RefreshRate     float64       `json:"refreshRate"`
	X               int           `json:"x"`
	Y               int           `json:"y"`
	ActiveWorkspace WorkspaceType `json:"activeWorkspace"`
	Reserved        []int         `json:"reserved"`
	Scale           float64       `json:"scale"`
	Transform       int           `json:"transform"`
	Focused         bool          `json:"focused"`
	DpmsStatus      bool          `json:"dpmsStatus"`
}

type Client struct {
	Address        string    `json:"address"`
	At             []int     `json:"at"`
	Size           []int     `json:"size"`
	Workspace      Workspace `json:"workspace"`
	Floating       bool      `json:"floating"`
	Monitor        int       `json:"monitor"`
	Class          string    `json:"class"`
	Title          string    `json:"title"`
	Pid            int       `json:"pid"`
	Xwayland       bool      `json:"xwayland"`
	Pinned         bool      `json:"pinned"`
	Fullscreen     bool      `json:"fullscreen"`
	FullscreenMode int       `json:"fullscreenMode"`
}

type Window struct {
	Address   string        `json:"address"`
	At        []int         `json:"at"`
	Size      []int         `json:"size"`
	Workspace WorkspaceType `json:"workspace"`
	Floating  bool          `json:"floating"`
	Monitor   int           `json:"monitor"`
	Class     string        `json:"class"`
	Title     string        `json:"title"`
	Pid       int           `json:"pid"`
	Xwayland  bool          `json:"xwayland"`
}

type Layers map[string]Layer

type Layer struct {
	Levels map[int][]LayerField `json:"levels"`
}

type LayerField struct {
	Address   string `json:"address"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	W         int    `json:"w"`
	H         int    `json:"h"`
	Namespace string `json:"namespace"`
}

type Devices struct {
	Mice []struct {
		Address      string  `json:"address"`
		Name         string  `json:"name"`
		DefaultSpeed float64 `json:"defaultSpeed"`
	} `json:"mice"`
	Keyboards []struct {
		Address      string `json:"address"`
		Name         string `json:"name"`
		Rules        string `json:"rules"`
		Model        string `json:"model"`
		Layout       string `json:"layout"`
		Variant      string `json:"variant"`
		Options      string `json:"options"`
		ActiveKeymap string `json:"active_keymap"`
		Main         bool   `json:"main"`
	} `json:"keyboards"`
	Tablets  []interface{} `json:"tablets"`
	Touch    []interface{} `json:"touch"`
	Switches []struct {
		Address string `json:"address"`
		Name    string `json:"name"`
	} `json:"switches"`
}

type Version struct {
	Branch        string        `json:"branch"`
	Commit        string        `json:"commit"`
	Dirty         bool          `json:"dirty"`
	CommitMessage string        `json:"commit_message"`
	Flags         []interface{} `json:"flags"`
}

type CursorPos struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Bind struct {
	Locked          bool   `json:"locked"`
	Mouse           bool   `json:"mouse"`
	Release         bool   `json:"release"`
	Repeat          bool   `json:"repeat"`
  NoConsuming     bool   `json:"non_consuming"`
	ModMask         int    `json:"modmask"`
	Submap          string `json:"submap"`
	Key             string `json:"key"`
	KeyCode         int    `json:"keycode"`
	Dispatcher      string `json:"dispatcher"`
	Arg             string `json:"arg"`
}
