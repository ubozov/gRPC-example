# grpc-example

[![CircleCI](https://circleci.com/gh/ubozov/gRPC-example.svg?style=svg&circle-token=140fd6f55ed20bea132bdd6b7c841ff7253231bb)](https://circleci.com/gh/ubozov/gRPC-example)

## generate proto

```bash
cd proto
protoc -I . service.proto --go_out=plugins=grpc:. service.proto
```
