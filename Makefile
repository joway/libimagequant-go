build:
	go build

dep:
	env GO111MODULE=on go mod download
	git submodule update --remote
