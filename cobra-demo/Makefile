all: cobra-demo

cobra-demo: main.go $(shell find -name '*.go' -type f)
	go build

clean:
	rm -rf cobra-demo
