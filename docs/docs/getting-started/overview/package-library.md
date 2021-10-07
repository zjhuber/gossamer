---
layout: default
title: Package Library
permalink: /welcome/package-library/
---

Gossamer is inspired by Substrate, a highly modular, flexible, and customizeable framework for building blockchains and this inspiration is reflected by the structure of Gossamer's package library.

This document provides an overview of the packages that make up Gossamer - for more detailed information about each package, please refer to [Gossamer's official Go reference documentation](https://pkg.go.dev/github.com/ChainSafe/gossamer).

Gossamer packages can be categorized into **four package types**, which directly correspond to top-level directories in the Gossamer code base:

- **[cmd package](#cmd-package)**

    - `cmd/gossamer` - this package encapsulates Gossamer's command-line interface and serve as the entrypoint for users

- **[dot packages](#dot-packages)**

    - `dot/...` - packages in this directory implement various standards that are relevant to the Polkadot Network

- **[lib packages](#lib-packages)**

    - `lib/...` - this directory contains packages that encapsulate reusable logic for building blockchains and more

- **[SCALE package](#scale-package)**

    - `pkg/scale` - this package provides [SCALE](https://substrate.dev/docs/en/knowledgebase/advanced/codec) encoding and decoding capabilities

- **[chain packages](#chain-packages)**

    - `chain/...` - the subdirectories in this directory contain configuration parameters for the blockchain networks that Gossamer supports

## cmd package

#### `cmd/gossamer`

- The entrypoint of this package is [`main.go`](https://github.com/ChainSafe/gossamer/blob/development/cmd/gossamer/main.go); this file defines the subcommands that the Gossamer command line interface accepts and implements "actions" that correspond to each of this subcommands. Gossamer uses [the popular `cli` package](https://github.com/urfave/cli/blob/master/docs/v1/manual.md) from `urfave` to provide its command line interface capabilities. The actions defined in `main.go` are:

    - `gossamerAction` - This is the action that is invoked when the Gossamer command line interface is executed without a subcommand; it accepts a number of parameters and will launch a Gossamer blockchain client.

## dot packages

#### `dot`

- The **dot package** contains packages that implement the Polkadot Host spec. The **dot package** implements the [Host Node](/getting-started/overview/host-architecture#host-node); it is the base node implementation for all [Official Nodes](/getting-started/overview/host-architecture#official-nodes) and [Custom Services](/getting-started/overview/host-architecture#custom-services) built with Gossamer.

#### `dot/core`

- The **core package** implements the [Core Service](/getting-started/overview/host-architecture#core-service) -  responsible for block production and block finalisation (consensus) and processing messages received from the [Network Service](/getting-started/overview/host-architecture/#network-service).

#### `dot/network`

- The **network package** implements the [Network Service](/getting-started/overview/host-architecture/#network-service) - responsible for coordinating network host and peer interactions. It manages peer connections, receives and parses messages from connected peers and handles each message based on its type.

#### `dot/state`

- The **state package** implements the [State Service](/getting-started/overview/host-architecture#state-service) - the source of truth for all chain and node state that is made accessible to all node services.

#### `dot/sync`

- The **sync package** implements handling of blocks received from the network layer.

#### `dot/rpc`

- The **rpc package** implements the [RPC Service](/getting-started/overview/host-architecture#rpc-service) - an HTTP server that handles state interactions.

#### `dot/types`

- The **types package** implements types defined within the Polkadot Host specification.

## lib packages

#### `lib/babe`

- the **babe package** implements the BABE block production algorithm.

#### `lib/blocktree`

- the **blocktree package** implements the blocktree, a data structure which tracks the chain and all its non-finalised forks.

#### `lib/common`

- the **common package** defines common types and functions.

#### `lib/crypto`

- the **crypto package** contains the key types used by the node (sr25519, ed25519, secp256k1).

#### `lib/grandpa`

- the **grandpa package** implements the GRANDPA finality gadget.

#### `lib/keystore`

- the **keystore package** manages the keystore and includes test keyrings.

#### `lib/runtime`

- the **runtime package** contains various wasm interpreters used to interpret the runtime. It currently contains `life`, `wasmer`, and `wasmtime`; however, `wasmer` is the only interpreter that is fully supported at the moment. In the future, all interpreters will be fully supported.

#### `lib/scale`

- the **scale package** implements the SCALE codec.

#### `lib/services`

- the **services package** implements a common interface for node services.

#### `lib/transaction`

- the **transaction package** is contains transaction types and the transaction queue data structure.

#### `lib/trie`

- the **trie package** implements a modified merkle-patricia trie.

#### `lib/utils`

- the **utils package** is used to manage node and test directories.
