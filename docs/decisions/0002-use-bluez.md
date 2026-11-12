# 0002. Use BlueZ, through its D-Bus API, as the Bluetooth stack

- **Status:** Accepted
- **Date:** 2026-11-08

## Context

The bridge needs fine control over Bluetooth: connecting to devices as a central, acting
as a peripheral toward the phone (advertising and serving a GATT database), forwarding
every GATT operation, pairing, and managing the adapter's identity.

Options considered:

- **BlueZ**, the standard Linux Bluetooth stack. Programs control it through its D-Bus
  API.
- **Our own Bluetooth host stack**, talking directly to the adapter over HCI. It gives
  full control, but it's a huge amount of work (pairing, encryption, L2CAP, ATT…)
  before even starting on the bridge.

In Go, there are two ways to use BlueZ:

- **`tinygo.org/x/bluetooth`**, a high-level library that talks to BlueZ over D-Bus on
  Linux. It's easy to start with, but it only supports BLE and exposes only part of what
  BlueZ offers.
- **`github.com/godbus/dbus/v5`**, to call the BlueZ D-Bus API directly. It takes more
  code, but gives full control.

## Decision

Use BlueZ through its D-Bus API. Use `tinygo.org/x/bluetooth` where it covers what we
need, such as scanning, and call BlueZ directly with `github.com/godbus/dbus/v5` where
it doesn't.

## Consequences

- BlueZ is mature and already handles pairing, key storage and the ATT protocol.
- Both libraries are pure Go on Linux, so gattroam builds as a single static binary
  without cgo and cross-compiles easily, for example to a Raspberry Pi.
- BlueZ, together with PipeWire, also supports Bluetooth Classic audio for later.
- gattroam only runs on Linux (see [0003](0003-linux-only.md)).
- BlueZ manages some things itself, such as the GAP and GATT services and where bonding
  keys are stored. The bridge has to work around them.
- Behavior can change between BlueZ versions, so bug reports must include the BlueZ
  version.
- The code talks to BlueZ in two different ways. Both must stay behind the same internal
  interface.
