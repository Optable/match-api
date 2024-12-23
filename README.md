# match-api
A set of public facing APIs defined by protobuf messages to perform a secure
PAIR (Publisher Advertiser Identity Reconciliation) match and activation with
Optable Data connectivity nodes (DCN). The open sourced Command Line Interface
(CLI) utility [opair](https://github.com/Optable/optable-pair-cli) uses
these APIs to enable any users of the tool to perform secure ID matches with a
DCN partner.

## Installation
Download Google's Protocol Buffers compiler `protoc` and make sure it's
available in your PATH. You can download it
[here](https://developers.google.com/protocol-buffers/docs/downloads), or
follow the installation instruction
[here](https://grpc.io/docs/protoc-installation/).

## Build
To compile and regenerate the go output files, simply run
```bash
make
```
