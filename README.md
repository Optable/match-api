# API
This repository contains the interface definitions (protobufs) of public Optable APIs for interacting with Optable Data connectivity platform to perform PSI. 

## Installation
Download Google's Protocol Buffers compiler `protoc` and make sure it's available in your PATH. You can download it
[here](https://developers.google.com/protocol-buffers/docs/downloads), or follow the installation instruction
[here](https://grpc.io/docs/protoc-installation/).

## Build
To regenerate the go protobuf files, simply run
```bash
make
```
