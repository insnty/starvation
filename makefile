BIN := starvation

.PHONY: build
build:
	go build -o $(BIN) main.go

.PHONY: clean
clean:
	rm -f $(BIN)