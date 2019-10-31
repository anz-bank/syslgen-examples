.PHONY: all
all: clean/todos mkdir/todos generate/todos gotools/todos compile

.PHONY: clean
clean: clean/todos

.PHONY: mkdir
mkdir: mkdir/todos

.PHONY: gotools
gotools: gotools/todos

.PHONY: generate
generate: generate/todos

.PHONY: validate
validate: validate/todos

.PHONY: compile
compile: todos_client/todos todos_server/todos

clean/%:
	-rm $*_client/$*-client
	-rm $*_server/$*-server

todos_client/%:
	cd $*_client; go build -o $*-client; cd -

todos_server/%:
	cd $*_server; go build -o $*-server; cd -

mkdir/%:
	echo "creating output dir" $*
	-mkdir -p $*

gotools/%:
	gofmt -w $*/
	goimports -w $*/
	golangci-lint run $*

generate/%:
	cd codegen && ./generate.sh