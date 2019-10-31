# SYSLGEN Examples

Examples of rest api code generation using `syslgen`.

See `examples/todos.sysl` for the rest api description that has been created for the following service : `http://jsonplaceholder.typicode.com/`

## How to build code

### Get syslgen binary

`syslgen` is expected to be available as a download in near future.

Currently the only way to get syslgen binary is by building it like so:

Assumes you have Go installed.
```
export GOPATH=$HOME/gopath
export SYSLBASE=$HOME/gopath/src/github.com/anz-bank
mkdir -p $SYSLBASE
cd $SYSLBASE
git clone https://github.com/anz-bank/sysl/
cd $SYSLBASE/sysl/sysl2
go get -t -v github.com/anz-bank/sysl/sysl2/sysl
go install -v
ln -s $GOPATH/bin/sysl syslgen
```
### Build code

Checkout the code and run the following commands in the root of the repository.

```bash
make
```
The make command will generate `todos/service.go` and build the binary `todos/todos`

Try to run the example like so:

```
./todos/todos
```

This client calls the generated service method `GET_todos_id`.


## File Organization

