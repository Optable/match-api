# Match-API
This repository contains the interface definitions (protobufs) of public Optable match APIs for performing secure Private Set Intersection (PSI) matches with Optable Data connectivity platforms (DCN).

## Installation
Download Google's Protocol Buffers compiler `protoc` and make sure it's available in your PATH. You can download it
[here](https://developers.google.com/protocol-buffers/docs/downloads), or follow the installation instruction
[here](https://grpc.io/docs/protoc-installation/).

## Build
To regenerate the go protobuf files, simply run
```bash
make
```
