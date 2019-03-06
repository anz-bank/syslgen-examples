.PHONY: all
all: todos apiconnect

SYSLGEN=$(PWD)/syslgen
REST_TRANSFORM = transforms/svc_interface.gen.sysl
TYPES_TRANSFORM = transforms/svc_types.gen.sysl
GRAMMAR = grammars/go.gen.g

REST_TRANSFORM_INPUT = $(REST_TRANSFORM) $(GRAMMAR)
TYPES_TRANSFORM_INPUT = $(TYPES_TRANSFORM) $(GRAMMAR)

todos: todos/todos
apiconnect: apiconnect/apiconnect

todos/todos: todos/service.go
	-rm todos/todos
	cd todos; go build -o todos; cd -

apiconnect/apiconnect: apiconnect/service.go
	-rm apiconnect/apiconnect
	cd apiconnect; go build -o apiconnect; cd -

%/service.go: examples/%.sysl $(TYPES_TRANSFORM_INPUT) %/main.go
	echo "creating output dir" $*
	-mkdir $*
	$(SYSLGEN) -root-model . -root-transform . -transform $(TYPES_TRANSFORM) -model examples/$*.sysl -grammar $(GRAMMAR) -start goFile -outdir $*
	gofmt -w $*/service.go
