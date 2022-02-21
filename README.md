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

The Flare binary, named `flare`, is in the `build` directory.

## Running Flare

### Legacy Version Upgrade

**If your node was previously running on the legacy version from the Gitlab repository, some directories need to be renamed/moved.**

- The default directory changed from `$HOME/.avalanchego` to `$HOME/.flare`.
- The name of the database sub-directory changed from `db/fuji` to `db/songbird` / `db/coston` respectively, to reflect the actual name of the network.

However, if you are running on the legacy version with default parameters, you are probably using RocksDB as the database engine.

**We highly recommend node operators to resynchronize their nodes using LevelDB as database engine.**

The RocksDB library used by the Avalanche code base is flawed and the database itself is a lot less reliable, thus being more liable to corruption.

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

These are the default settings:

```json
{
  "snowman-api-enabled": false,
  "coreth-admin-api-enabled": false,
  "coreth-admin-api-dir": "",
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
  "continuous-profiler-dir": "",
  "continuous-profiler-frequency": 900000000000,
  "continuous-profiler-max-files": 5,
  "rpc-gas-cap": 50000000,
  "rpc-tx-fee-cap": 100,
  "preimages-enabled": false,
  "pruning-enabled": true,
  "snapshot-async": true,
  "snapshot-verification-enabled": false,
  "metrics-enabled": false,
  "metrics-expensive-enabled": false,
  "local-txs-enabled": false,
  "api-max-duration": 0,
  "ws-cpu-refill-rate": 0,
  "ws-cpu-max-stored": 0,
  "api-max-blocks-per-request": 0,
  "allow-unfinalized-queries": false,
  "allow-unprotected-txs": false,
  "keystore-directory": "",
  "keystore-external-signer": "",
  "keystore-insecure-unlock-allowed": false,
  "remote-tx-gossip-only-enabled": false,
  "tx-regossip-frequency": 60000000000,
  "tx-regossip-max-size": 15,
  "log-level": "debug",
  "offline-pruning-enabled": false,
  "offline-pruning-bloom-filter-size": 512,
  "offline-pruning-data-directory": ""
}
```

You can refer to the original [Avalanche documentation](https://docs.avax.network/build/references/avalanchego-config-flags/#c-chain-configs) for a description of the settings.

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
