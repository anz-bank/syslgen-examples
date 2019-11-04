.PHONY: all
all: clean/bin generate/todos gotools/todos compile

.PHONY: clean
clean: clean/bin

.PHONY: gotools
gotools: gotools/todos

.PHONY: generate
generate: generate/todos

.PHONY: validate
validate: validate/todos

.PHONY: compile
compile: todos_client todos_server

clean/%:
	-rm $*_client/$*-client
	-rm $*_server/$*-server

todos_client:
	cd cmd/todos_client; go build -o ../../bin/client *.go; cd -

todos_server:
	cd cmd/todos_server; go build -o ../../bin/server *.go; cd -

mkdir/%:
	echo "creating output dir" $*
	-mkdir -p $*

gotools/%:
	gofmt -w $*/
	goimports -w $*/
	golangci-lint run $*

generate/%:
	cd codegen && ./generate.sh