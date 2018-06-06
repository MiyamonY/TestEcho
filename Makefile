TARGET := TestEcho
.PHONY: install clean

server:
	go run main.go

build: main.go
	go build -o $(TARGET)

install:
	dep ensure
	npm install
	cd assets/semantic && gulp build

clean:
	$(RM) -rf node_modules assets/semantic $(TARGET)
