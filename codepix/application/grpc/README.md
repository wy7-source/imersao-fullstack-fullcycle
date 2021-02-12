# Imersão Full Stack & FullCycle - CodePix
Arquivo com o comando utilizado para gerar os Stubs do gRPC.

```
protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto
```
