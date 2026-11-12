# 0003. Support Linux only

- **Status:** Accepted
- **Date:** 2026-11-08

## Context

gattroam should run on any machine with a Bluetooth adapter. But the bridge needs a level
of control that most operating systems don't give to programs. Through their standard
APIs, Windows and macOS don't let a program change the adapter's address, and macOS
doesn't let a program act as headphones (an A2DP sink), which Bluetooth Classic audio
will need.

[ADR 0002](0002-use-bluez.md) also chose BlueZ, which only exists on Linux.

## Decision

gattroam supports Linux only. An access point is any Linux machine with a Bluetooth 4.0+
adapter, including small boards such as a Raspberry Pi.

## Consequences

- There is only one platform to develop and test.
- Cheap hardware, such as a Raspberry Pi, works as an access point.
- Nothing is installed on phones or Bluetooth devices, so this decision doesn't limit
  which phones or devices can use gattroam.
- People who only have Windows or macOS machines can't run an access point.
- Worth revisiting if there is demand and those platforms' Bluetooth APIs allow it.
