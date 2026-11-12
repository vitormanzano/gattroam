# 0001. Start with BLE (GATT) before Bluetooth Classic

- **Status:** Accepted
- **Date:** 2026-11-08

## Context

gattroam's goal is to bridge any Bluetooth device across access points. But "Bluetooth"
is really two technologies:

- **Bluetooth Low Energy (BLE)**, used by fitness bands, sensors, tags and many other
  small devices. Almost every BLE device exchanges data through GATT, which works the
  same way on every device.
- **Bluetooth Classic**, used by headphones, speakers and car kits. Each use has its own
  profile (A2DP for audio, HFP for calls, AVRCP for media controls, HID for input
  devices, and others), and each profile needs its own bridging code, with its own timing
  constraints.

Handoff between access points also differs. For a device to move between access points
without noticing, all of them must present the same Bluetooth identity. In BLE, an
adapter can use a random static address as its identity, set with a standard command
that every BLE adapter supports. In Bluetooth Classic, the identity is the adapter's
public address, which can only be changed with vendor-specific commands that many
adapters don't support.

## Decision

The first version of gattroam bridges BLE devices through a single, generic GATT bridge.
Bluetooth Classic, starting with audio (headphones and speakers), comes later.

## Consequences

- One generic implementation covers a wide range of devices.
- Seamless handoff is achievable on any BLE adapter.
- The project name, gattroam, matches what it does.
- Headphones and speakers, the use case many people think of first, aren't supported in
  the first version.
- This decision doesn't cover LE Audio either: although it runs over BLE, the audio
  travels over isochronous channels, not GATT.
