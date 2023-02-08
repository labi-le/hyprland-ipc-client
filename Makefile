PROJ_NAME = hyprland-ipc

tests:
	go test -coverpkg=./... ./... -parallel=2