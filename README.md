# SYSLGEN Examples

This is an example repo showing the generation of REST API client and server code using `syslgen`.
The example application consists an API Server that stores tasks, and a client CLI tool that makes API calls to the server.

See [codegen/model/todos.sysl](codegen/model/todos.sysl) for the REST API description

## Getting Started

### Get Syslgen

#### Download Binary

`syslgenv0.2.8` can be downloaded as a binary from [https://github.com/anz-bank/sysl/releases/tag/v0.2.8](https://github.com/anz-bank/sysl/releases/tag/v0.2.8)

You can then add it to your PATH using

```bash
sudo ln -s <absolute-path-to-gosysl-binary> /usr/local/bin/syslgen
```

#### Build from Source

To get the latest version, you can also download and build the binary like so:
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

### Build Example Code

Checkout the code and run the following commands in the root of the repository.

```bash
make
```

The make command will generate the following files

- [./gen/todos/service.go](./gen/todos/service.go)
- [./gen/todos/requestrouter.go](./gen/todos/requestrouter.go)
- [./gen/todos/servicehandler.go](./gen/todos/servicehandler.go)
- [./gen/todos/serviceinterface.go](./gen/todos/serviceinterface.go)
- [./gen/todos/types.go](./gen/todos/types.go)

and build the binaries `bin/client` and `bin/server`

### Running the example applications

Try to run the server like so:

``` bash
./bin/server
```

We can then use the client to retrieve all posts:

``` bash
./bin/client posts
```

We can also create a post:

``` bash
./bin/client comment test1234 test
```

For more information try the following:

``` bash
./bin/client --help
```

## Developing

### Grammars

Grammars define what code looks like in various languages.

The grammar definition for sysl is in [codegen/grammars/go.gen.sysl](codegen/grammars/go.gen.sysl)
The sysl grammar defines types used in the transform files.

The grammar definition for go is in [codegen/grammars/go.gen.g](codegen/grammars/go.gen.g)
This defines what the output files should look like.

### Model

This is a description of the application's endpoints and data types

### Transforms

For each file that we want to generate, we need a transform.

Transform files start with a toplevel declaration of CodeGenTransform followed by a number of view declarations.

```
CodeGenTransform:
  !view filename(app <: sysl.App) -> string:
    app -> (:
      filename =  "service.go"
    )
```

A view declaration is similar to a virtual table, and can be likened to a function. The key ones to pay attention to are:

- goFile View
  - This view controls what goes into the generated go file
  - Can think of this as the main function