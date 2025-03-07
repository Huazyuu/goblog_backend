.PHONY: build run gotool clean help

BINARY="./gvb_exe/gvb"

build :
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY}.exe

buildwin :
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY}.exe

buildlinux :
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run : 
	go run main.go

db :
	go run main.go -db

gotool :
	go fmt ./
	go vet ./

es create:
	go run main.go -es create

es remove:
	go run main.go -es remove

swag :
	swag init
clean :
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
air init:
	air init
air run:
	air run 

help :
	@echo "make build [-linux/windows] - Compile the Go code and generate the binary"
	@echo "make run - run main.go"
	@echo "make init db - migrate database"
	@echo "make clean - remove binary code files and vim swap files"
	@echo "make gotool - run Go tools : 'fmt' and 'vet'"
	@echo "make swag - generate api doc"
	@echo "make air - project automatically restarts"

