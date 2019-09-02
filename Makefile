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

SYSLGEN=$(GOPATH)/bin/syslgen
TYPES_TRANSFORM = transforms/svc_types.sysl
CLIENT_TRANSFORM = transforms/svc_client.sysl
INTERFACE_TRANSFORM = transforms/svc_interface.sysl
HANDLER_TRANSFORM = transforms/svc_handler.sysl
ROUTER_TRANSFORM = transforms/svc_router.sysl
GRAMMAR = grammars/go.gen.g
TYPES_TRANSFORM_INPUT = $(TYPES_TRANSFORM) $(GRAMMAR)

gen = $(SYSLGEN) gen --root-model . --root-transform . --transform $(1) --model examples/$(2).sysl --grammar $(GRAMMAR) --start goFile --outdir $(2)
validate = $(SYSLGEN) validate --grammar $(GRAMMAR) --start goFile --root-transform . --transform $(1)

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
	$(call gen,$(TYPES_TRANSFORM),$*)
	$(call gen,$(CLIENT_TRANSFORM),$*)
	$(call gen,$(INTERFACE_TRANSFORM),$*)
	$(call gen,$(HANDLER_TRANSFORM),$*)
	$(call gen,$(ROUTER_TRANSFORM),$*)

validate/%:
	$(call validate,$(TYPES_TRANSFORM),$*)
	$(call validate,$(CLIENT_TRANSFORM),$*)
	$(call validate,$(INTERFACE_TRANSFORM),$*)
	$(call validate,$(HANDLER_TRANSFORM),$*)
	$(call validate,$(ROUTER_TRANSFORM),$*)
