# gattroam

Keep your Bluetooth devices connected across your access points, no matter how far they roam.

> **Status:** early development — nothing works yet. Ideas and contributions are welcome.

## The problem

Bluetooth only works while both devices are within radio range of each other, usually
around 10 meters indoors. Walk away from your phone, and whatever is connected to it
(a fitness band, a sensor, your headphones) disconnects.

## How gattroam works

You run gattroam on access points spread around your home or office. Each access point
acts as a bridge between your phone and your Bluetooth device:

- to your **phone**, gattroam looks like the original device;
- to the **device**, gattroam looks like your phone.

Data goes from your phone to the access point closest to it, travels over the network to
the access point closest to the device, and from there to the device. When the device
moves, the access point nearest to it takes over the connection.

```
phone ──Bluetooth──▶ [access point A] ══network══▶ [access point B] ──Bluetooth──▶ device
```

All access points share the same Bluetooth identity, so a device that moves to a new
access point reconnects as if nothing had changed: your phone only notices a short delay.

Nothing needs to be installed on the phone or the device; they only see a regular
Bluetooth device. The only thing you install is the gattroam binary on each access point.

## Supported devices

The first version targets **Bluetooth Low Energy (BLE)** devices that use GATT, such as
fitness bands, heart-rate monitors, sensors and tags. GATT works the same way on every
BLE device, so a single generic bridge covers all of them.

**Bluetooth Classic** devices, such as headphones and speakers, are planned for later
(see the [roadmap](#roadmap)).

Some devices are designed to detect and block exactly this kind of bridge, for example
digital car keys that measure distance or latency. gattroam will not work with them.

## Requirements

- Linux with BlueZ (the Linux Bluetooth stack).
- A Bluetooth 4.0+ adapter, built-in or USB, that can keep two connections at the same
  time: one to your phone and one to the device.
- All access points on the same network.

An "access point" is any Linux computer with Bluetooth, such as a Raspberry Pi. Most
Wi-Fi routers don't have Bluetooth, so they can't run gattroam.

## Expected limitations

- **Latency:** every extra hop adds delay. That's usually fine for sensors, but it may be
  noticeable with keyboards, game controllers and, later, audio.
- **Pairing:** your phone and your device pair with gattroam instead of with each other.
  Because all access points share the same identity, you only pair once.

## Security

gattroam relays Bluetooth traffic over your network, which is the same technique used in
Bluetooth relay attacks. It is meant for your own devices: gattroam only bridges devices
you explicitly pair with it, and traffic between access points will be authenticated and
encrypted.

To report a vulnerability, see [SECURITY.md](SECURITY.md).

## Roadmap

- [ ] Bridge a BLE device to a phone through a single access point
- [ ] Bridge through two access points over the network
- [ ] Authenticated, encrypted traffic between access points
- [ ] Handoff between access points
- [ ] Bluetooth Classic audio (headphones and speakers)

## Contributing

gattroam is in its early stages, and ideas are welcome. Before getting started, read the
[contributing guide](CONTRIBUTING.md) and the [code of conduct](CODE_OF_CONDUCT.md).
The important design decisions are recorded in [docs/decisions](docs/decisions/README.md).

## License

[MIT](LICENSE)
