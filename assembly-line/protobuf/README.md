# protobuf

## generate code

1. cd to [pb](./pb)

```bash
$ cd ./pb
```

2. generate code

```bash
$ protoc --proto_path=. --go_out=paths=source_relative:../go  cpu/cpu.proto --micro_out=paths=source_relative:../go
```