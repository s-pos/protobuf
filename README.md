## Protobuf

### 1. Installation
- For Windows user using WSL
> ```bash
> $ sudo apt install -y protobuf-compiler
> $ protoc --version
> ```

- For MacOS user
> ```bash
> $ brew install protobuf
> $ protoc --version
> ```

### 2. Extension
- Install Protoc GO Extension (gRPC) with the syntax
> with Go version >= 1.16
> ```bash
> $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
> $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
> ```
After install Extension, you need export `PATH` binary of Go with syntax
```bash
export PATH=$PATH:$GOPATH/bin
```

### 3. Generate
For generate file proto, please run this syntax
```bash
$ protoc \
    --go_out=./go \
    --go-grpc_out=./go \
    ./pb/*/*.proto
```