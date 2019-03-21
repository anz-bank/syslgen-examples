.PHONY: all
all: todos

SYSLGEN=$(PWD)/syslgen
TYPES_TRANSFORM = transforms/svc_types.gen.sysl
GRAMMAR = grammars/go.gen.g
TYPES_TRANSFORM_INPUT = $(TYPES_TRANSFORM) $(GRAMMAR)

todos: todos_client/todos

todos_client/%: %/service.go %_client/main.go
	-rm $*/$*
	cd $*_client; go build -o $*; cd -

%/service.go: examples/%.sysl $(TYPES_TRANSFORM_INPUT)
	echo "creating output dir" $*
	-mkdir $*
	$(SYSLGEN) -root-model . -root-transform . -transform $(TYPES_TRANSFORM) -model examples/$*.sysl -grammar $(GRAMMAR) -start goFile -outdir $*
	gofmt -w $*/service.go
	golangci-lint run $*
