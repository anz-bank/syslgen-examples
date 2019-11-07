#!/bin/sh

set -e -x

SYSLGEN="syslgen"
TRANSFORMS="transforms"
GRAMMAR="grammars/go.gen.g"
START="goFile"

TODOS_MODEL="model/todos.sysl"
TODOS_OUT="../gen/todos"

declare -a allxforms=(svc_client.sysl svc_types.sysl svc_router.sysl svc_handler.sysl svc_interface.sysl)

for xform in "${allxforms[@]}"
do
    syslgen gen --root . --root-transform . --transform $TRANSFORMS/$xform --grammar $GRAMMAR --start $START --outdir $TODOS_OUT --app-name Todos $TODOS_MODEL
done

go fmt $TODOS_OUT/*.go
goimports -w $TODOS_OUT/*.go

