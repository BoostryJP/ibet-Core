# Repository Guidelines

## Project Structure & Module Organization

ibet-Core is the ibet Network node client, independently maintained from a GoQuorum fork. The Go module retains `github.com/ethereum/go-ethereum` as its import path.

- `cmd/`: executable entry points, including `geth`, `bootnode`, and `abigen`.
- `core/`, `consensus/`, `eth/`, `p2p/`, and `rpc/`: execution, consensus, networking, and APIs.
- `private/`, `permission/`, `plugin/`, and `raft/`: inherited enterprise and permissioned-network features.
- Package-local `*_test.go` files and `testdata/` directories contain tests and fixtures; `tests/` holds protocol tests and fuzzers.
- `contracts/` contains contract sources; `docs/` contains operational documentation. `build/` contains build tooling; generated executables go to `build/bin/`.

## Build, Test, and Development Commands

Install Go 1.26.4 (matching `.go-version` and `go.mod`) and a C compiler. Run commands from the repository root.

- `git submodule update --init --recursive`: fetch test fixtures under `tests/testdata`.
- `make geth`: build the node as `build/bin/geth`.
- `make all`: build all command-line utilities.
- `./build/bin/geth --dev --datadir /tmp/ibet-core-dev console`: launch a local development node and console.
- `make test`: build utilities, then run Go tests with package execution serialized.
- `go test ./core/...`: run focused package tests.
- `make format`: format Go source.
- `make lint`: run the pinned golangci-lint tool using `.golangci.yml`.

## Coding Style & Naming Conventions

Use standard Go formatting with tab indentation and goimports-compatible imports. Keep package names lowercase, exported identifiers in PascalCase, and unexported identifiers in camelCase. Follow nearby code conventions and preserve upstream license headers. Update generators rather than hand-editing generated files.

## Testing Guidelines

Use Go's `testing` framework, naming files `*_test.go` and functions `TestXxx` or `BenchmarkXxx`. Add regression tests for behavior changes and keep fixtures in `testdata/`. Run focused tests during development and `make test` and `make lint` before submitting code. Use `go run build/ci.go test -race ./path/to/package` for concurrency changes. No numeric coverage threshold is configured.

## Commit & Pull Request Guidelines

Recent commits use concise imperative subjects such as `Fix Prague/Osaka precompile selection`; dependency updates use `chore(go)(deps): ...` or `chore(ci)(deps): ...`. Keep commits focused. PR descriptions should explain the problem, resulting behavior, related issues, and validation commands. Highlight consensus, fork-activation, configuration, or compatibility changes that affect node operators.
