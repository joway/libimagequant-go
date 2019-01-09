build:
	go build

dep:
	env GO111MODULE=on go mod download
	git submodule update --remote
	cd lib && make static && ./configure --prefix=/usr && make libimagequant.so && cd ..
