.PHONY: all
all: clean generate compile

# Generate doesn't run automatically on make all. It can be run manually via make generate
.PHONY: generate
generate: generate/todos

.PHONY: compile
compile: todos_client todos_server

.PHONY: clean
clean:
	-rm bin/client
	-rm bin/server

.PHONY: todos_client
todos_client:
	go build -o ./bin/client cmd/todos_client/*.go

.PHONY: todos_server
todos_server:
	go build -o ./bin/server cmd/todos_server/*.go

generate/%:
	cd codegen && ./generate.sh