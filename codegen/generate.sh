#!/bin/sh

set -e -x

TRANSFORMS="transforms"
GRAMMAR="grammars/go.gen.g"
START="goFile"
GOFLAGS="-mod=vendor"

TODOS_MODEL="model/todos.sysl"
TODOS_OUT="../gen/todos"

declare -a clientxforms=(svc_client.sysl svc_types.sysl)
declare -a serverxforms=(svc_types.sysl svc_router.sysl svc_handler.sysl svc_interface.sysl)
declare -a allxforms=(svc_client.sysl svc_types.sysl svc_router.sysl svc_handler.sysl svc_interface.sysl)

for xform in "${allxforms[@]}"
do
    syslgen gen --root . --root-transform . --transform $TRANSFORMS/$xform --grammar $GRAMMAR --start $START --outdir $TODOS_OUT --app-name Todos $TODOS_MODEL
done

declare -a outdirs=($TODOS_OUT)
for od in "${outdirs[@]}"
do
    go fmt $od/*.go
    goimports -w $od/*.go
done
