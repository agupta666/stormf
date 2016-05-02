
stormf: main.go
	go build

stormf.ubuntu:
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .

image: stormf.ubuntu
	docker build -t stormf .

clean:
	go clean

