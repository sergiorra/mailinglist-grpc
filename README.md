<div align="center">
  <h1>Mailing List gRPC API</h1>
  <blockquote>Simple gRPC API to manage a mailing list</blockquote>
</div>

<br/>

## 📜 Information

### Install protobuf compiler (Mac)

```
brew install protobuf
```

### Install Go protobuf codegen tools

`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

### Generate Go code from .proto files
After creating/updating your `.proto` file, this command will generate the necessary code.

```
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  Proto/mail.proto
```

## ⚙️ Usage
Execute each of the commands on separate command lines

```
go run ./server
go run ./client
```