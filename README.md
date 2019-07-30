# grpc-test

```bash
cd server
go build .
./server

cd client
go build .
./client
```

if you modified helloworld/helloworld.proto

```bash
protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
```
