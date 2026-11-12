# Contributing to gattroam

Thanks for your interest in gattroam! The project is in early development, so right now
the most useful contributions are ideas, design discussions and, soon, testing on
different hardware.

By participating, you agree to follow the [Code of Conduct](CODE_OF_CONDUCT.md).

## Ways to contribute

- **Ideas and design:** open an issue to discuss them.
- **Hardware testing:** once gattroam can scan for devices (v0.1), reports about which
  Bluetooth adapters and devices work, and which don't, will be very valuable.
- **Documentation:** fixes and clarifications are always welcome.
- **Code:** see below.

## Before you send a pull request

Please open an issue first and describe what you want to change, so we can agree on the
approach before you spend time on it. Small fixes, such as typos, can go straight to a
pull request.

Architectural changes, such as the Bluetooth stack, the network protocol between access
points or the supported platforms, need an Architecture Decision Record. See
[docs/decisions](docs/decisions/README.md).

## Development setup

You need:

- Go, at the version listed in [go.mod](go.mod);
- to run gattroam against real devices: Linux with BlueZ and a Bluetooth 4.0+ adapter.

Build, test and check the code:

```sh
go build -o gattroam ./cmd
go test ./...
go vet ./...
```

## Guidelines

- Format your code with `gofmt`.
- Unit tests must not need Bluetooth hardware. Keep Bluetooth access behind Go
  interfaces, so the bridge logic can be tested with fakes.
- Keep gattroam pure Go (no cgo), so it builds as a single static binary for any Linux
  machine, including a Raspberry Pi.
- Write code, comments, commit messages and documentation in English.
- Keep pull requests small and focused on a single change.

## Reporting bugs

Bluetooth bugs are hard to reproduce without details. Please include:

- what you did, what you expected and what happened instead;
- your Bluetooth adapter model, and whether it is built-in or USB;
- your Linux distribution and kernel version (`uname -r`);
- your BlueZ version (`bluetoothctl --version`);
- the device you tried to bridge;
- gattroam's logs and, if possible, a `btmon` capture of the problem.

`btmon` captures can contain device addresses and pairing keys. Remove them before
posting a capture publicly.

For security problems, **don't** open a public issue. See [SECURITY.md](SECURITY.md).
