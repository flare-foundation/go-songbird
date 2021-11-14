# Flare

Node implementation for the [Flare](https://flare.network) network.

## Installation

Flare is an incredibly lightweight protocol, so the minimum computer requirements are quite modest.
Note that as network usage increases, hardware requirements may change.

- CPU: Equivalent of 8 AWS vCPU
- RAM: 16 GB
- Storage: 200 GB
- OS: Ubuntu 18.04/20.04 or MacOS >= Catalina
- Network: Reliable IPv4 or IPv6 network connection, with an open public port.
- Software Dependencies:
  - [Go](https://golang.org/doc/install) version >= 1.16.8 and set up [`$GOPATH`](https://github.com/golang/go/wiki/SettingGOPATH).
  - [gcc](https://gcc.gnu.org/)
  - g++

### Native Install

Clone the Flare repository:

```sh
git clone https://github.com/flare-foundation/flare.git
```

#### Building the Flare Executable

Build Flare using the build script:

```sh
./scripts/build.sh
```

The Flare binary, named `flare`, is in the `build` directory.

## Running Flare

### Connecting to Songbird

To connect to the Songbird canary network, run:

```sh
./build/flare --network-id=songbird \
  --bootstrap-ips="$(curl -m 10 -sX POST --data '{ "jsonrpc":"2.0", "id":1, "method":"info.getNodeIP" }' -H 'content-type:application/json;' https://songbird.flare.network/ext/info | jq -r ".result.ip")" \
  --bootstrap-ids="$(curl -m 10 -sX POST --data '{ "jsonrpc":"2.0", "id":1, "method":"info.getNodeID" }' -H 'content-type:application/json;' https://songbird.flare.network/ext/info | jq -r ".result.nodeID")"
```

You should see some _fire_ ASCII art and log messages.

You can use `Ctrl+C` to kill the node.

Please note that you currently need to be whitelisted to connect to beacon node.

### Launching Flare locally

To create a single node local test network, run:

```sh
./build/flare --network-id=local \
  --staking-enabled=false \
  --snow-sample-size=1 \
  --snow-quorum-size=1
```

This launches a Flare network with one node.

## Generating Code

Flare uses multiple tools to generate efficient and boilerplate code.

### Running protobuf codegen

To regenerate the protobuf go code, run `scripts/protobuf_codegen.sh` from the root of the repo.

This should only be necessary when upgrading protobuf versions or modifying .proto definition files.

To use this script, you must have [protoc](https://grpc.io/docs/protoc-installation/) (v3.17.3), protoc-gen-go (v1.26.0) and protoc-gen-go-grpc (v1.1.0) installed. protoc must be on your $PATH.

To install the protoc dependencies:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

If you have not already, you may need to add `$GOPATH/bin` to your `$PATH`:

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

If you extract protoc to ~/software/protobuf/, the following should work:

```sh
export PATH=$PATH:~/software/protobuf/bin/:~/go/bin
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/protobuf/cmd/protoc-gen-go-grpc
scripts/protobuf_codegen.sh
```

For more information, refer to the [GRPC Golang Quick Start Guide](https://grpc.io/docs/languages/go/quickstart/).

## Security Bugs

**We and our community welcome responsible disclosures.**

If you've discovered a security vulnerability, please report it via our [contact form](https://flare.network/contact/). Valid reports will be eligible for a reward (terms and conditions apply).
