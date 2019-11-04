
all: 
	go get
	GOBIN=`pwd`/bin go build main.go
