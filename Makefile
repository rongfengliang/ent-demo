.PHONY : clean
all: clean cli
clean: 
	rm -rf build/*
cli:
	go build -o build ./cmd/...
