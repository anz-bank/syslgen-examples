.PHONY: all
all: clean generate compile

# Generate doesn't run automatically on make all. It can be run manually via make generate
.PHONY: generate
generate: generate/todos

.PHONY: compile
compile: todosClient todosServer

.PHONY: clean
clean:
	-rm bin/client
	-rm bin/server
	-rm gen/todos/*.go

.PHONY: todosClient
todosClient:
	go build -o ./bin/client cmd/todosClient/*.go

.PHONY: todosServer
todosServer:
	go build -o ./bin/server cmd/todosServer/*.go

generate/%:
	cd codegen && ./generate.sh