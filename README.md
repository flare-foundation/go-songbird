# Flare

Node implementation for the [Flare](https://flare.network) network.

## Installation

Flare uses a relatively lightweight consensus protocol, so the minimum computer requirements are modest.
Note that as network usage increases, hardware requirements may change.

The minimum recommended hardware specification for nodes connected to Mainnet is:

- CPU: Equivalent of 8 AWS vCPU
- RAM: 16 GiB
- Storage: 512 GiB
- OS: Ubuntu 18.04/20.04 or macOS >= 10.15 (Catalina)
- Network: Reliable IPv4 or IPv6 network connection, with an open public port.

If you plan to build Flare from source, you will also need the following software:

- [Go](https://golang.org/doc/install) version >= 1.16.8
- [gcc](https://gcc.gnu.org/)
- g++

### Native Install

Clone the Flare repository:

```sh
git clone https://github.com/flare-foundation/flare.git
cd flare
```

This will clone and checkout to `master` branch.

### Building the Flare Executable

Build Flare using the build script:

```sh
./scripts/build.sh
```

If you want to build the binary with RocksDB support, you need to export the corresponding environment variable _before_ building:

```sh
export ROCKSDBALLOWED=1
```

The Flare binary, named `flare`, is in the `build` directory.


## Running Flare

### Legacy Version Upgrade

**Please note that the default database engine has changed from RocksDB to LevelDB.**

The RocksDB library used by the upstream Avalanche code base is flawed and the database itself is liable to data corruption.
LevelDB is a more mature and stable choice, and should provide a more reliable experience, while using less storage space.

**If you still want to use your previous database, there are a number of changes you need to be aware of:**

1. The default parent directory for logs, databases and configuration files changed from `$HOME/.avalanchego` to `$HOME/.flare`.
2. The name of the sub-directories in the database directory changed from `db/fuji` to `db/songbird` and `db/coston` respectively.
3. The node needs to be built with the `ROCKSDBALLOWED=1` environment variable and launched with the `--db-type=rocksdb` flag.

**However, we recommend that all node operators resync their nodes from scratch using LevelDB.**

### Connecting to Coston

To connect to the Coston test network, run:

```sh
./build/flare --network-id=coston \
  --bootstrap-ips="$(curl -m 10 -sX POST --data '{ "jsonrpc":"2.0", "id":1, "method":"info.getNodeIP" }' -H 'content-type:application/json;' https://coston.flare.network/ext/info | jq -r ".result.ip")" \
  --bootstrap-ids="$(curl -m 10 -sX POST --data '{ "jsonrpc":"2.0", "id":1, "method":"info.getNodeID" }' -H 'content-type:application/json;' https://coston.flare.network/ext/info | jq -r ".result.nodeID")"
```

You should see some _fire_ ASCII art and log messages.

You can use `Ctrl+C` to kill the node.

If you want your node's API to be reachable, you have to add the `--http-host=<ip_address>` flag to the command line.

### Connecting to Songbird

To connect to the Songbird canary network, run:

```sh
./build/flare --network-id=songbird \
  --bootstrap-ips="$(curl -m 10 -sX POST --data '{ "jsonrpc":"2.0", "id":1, "method":"info.getNodeIP" }' -H 'content-type:application/json;' https://songbird.flare.network/ext/info | jq -r ".result.ip")" \
  --bootstrap-ids="$(curl -m 10 -sX POST --data '{ "jsonrpc":"2.0", "id":1, "method":"info.getNodeID" }' -H 'content-type:application/json;' https://songbird.flare.network/ext/info | jq -r ".result.nodeID")"
```

You should see some _fire_ ASCII art and log messages.

You can use `Ctrl+C` to kill the node.

If you want your node's API to be reachable, you have to add the `--http-host=<ip_address>` flag to the command line.

Please note that you currently need to be whitelisted in order to connect to the Songbird network.

### Pruning & APIs

The configuration for the chain is loaded from a configuration file, located at `{chain-config-dir}/C/config.json`.

Here are the most relevant default settings:

```json
{
  "snowman-api-enabled": false,
  "coreth-admin-api-enabled": false,
  "eth-apis": [
    "public-eth",
    "public-eth-filter",
    "net",
    "web3",
    "internal-public-eth",
    "internal-public-blockchain",
    "internal-public-transaction-pool",
    "internal-public-account"
  ],
  "rpc-gas-cap": 50000000,
  "rpc-tx-fee-cap": 100,
  "pruning-enabled": true,
  "local-txs-enabled": false,
  "api-max-duration": 0,
  "api-max-blocks-per-request": 0,
  "allow-unfinalized-queries": false,
  "allow-unprotected-txs": false,
  "remote-tx-gossip-only-enabled": false,
  "log-level": "debug",
}
```

You can refer to the original [Avalanche documentation](https://docs.avax.network/build/references/avalanchego-config-flags/#c-chain-configs) for a full list of all settings and a detailed description.

The directory for configuration files defaults to `HOME/.flare/configs` and can be changed using the `--chain-config-dir` flag.

In order to disable pruning and run a full archival node, `pruning-enabled` should be set to `false`.

The various node APIs can also be enabled and disabled by setting the respective parameters.

### Launching Flare locally

In order to run a local network, the validator set needs to be defined locally.
This can be done by setting the validator set in a environment variable.

You can use `./scripts/launch_localnet.sh` as an easy way to spin up a 5-node local network.
All funds are controlled by the private key under `/.scripts/keys/6b0dd034a2fd67b932f10e3dba1d2bbd39348695.json`.

## Generating Code

Flare uses multiple tools to generate boilerplate code.

### Running protobuf codegen

To regenerate the protobuf go code, run `scripts/protobuf_codegen.sh` from the root of the repo.

This should only be necessary when upgrading protobuf versions or modifying .proto definition files.

To use this script, you must have [buf](https://docs.buf.build/installation) (v1.0.0-rc12), protoc-gen-go (v1.27.1) and protoc-gen-go-grpc (v1.2.0) installed.

To install the buf dependencies:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
```

If you have not already, you may need to add `$GOPATH/bin` to your `$PATH`:

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

If you extract buf to ~/software/buf/bin, the following should work:

```sh
export PATH=$PATH:~/software/buf/bin/:~/go/bin
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/protobuf/cmd/protoc-gen-go-grpc
scripts/protobuf_codegen.sh
```

For more information, refer to the [GRPC Golang Quick Start Guide](https://grpc.io/docs/languages/go/quickstart/).        |

## Security Bugs

**We and our community welcome responsible disclosures.**

If you've discovered a security vulnerability, please report it via our [contact form](https://flare.network/contact/). Valid reports will be eligible for a reward (terms and conditions apply).
