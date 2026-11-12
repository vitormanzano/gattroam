# Security Policy

gattroam relays Bluetooth traffic over a network, which is the same technique used in
Bluetooth relay attacks. Security is a core part of the project, not an add-on.

## Supported versions

gattroam is in early development and has no releases yet. Once there are releases,
security fixes will go to the latest release only.

## Reporting a vulnerability

Please **don't** open a public issue for security problems.

Report them privately through GitHub instead:
[report a vulnerability](https://github.com/vitormanzano/gattroam/security/advisories/new).
Include a description of the problem, the steps to reproduce it, its impact, and the
version or commit you tested.

gattroam is a volunteer project, so response times may vary, but every report will be
read and taken seriously. Once a problem is fixed, it will be published in a security
advisory, with credit to you if you want it.

## Threat model

> This is the planned design. gattroam doesn't implement these protections yet; they are
> tracked in the [roadmap](README.md#roadmap).

gattroam aims to protect against:

- **Unauthorized access points:** only nodes you explicitly authorize can join your
  gattroam network.
- **Eavesdropping and tampering on the network:** traffic between access points is
  authenticated and encrypted.
- **Bridging devices without consent:** gattroam only bridges devices you explicitly pair
  with it.

Out of scope:

- **An attacker with root access to an access point.** They can read its keys and
  traffic anyway.
- **Weaknesses in the Bluetooth link itself.** The link between an access point and a
  device is only as secure as that device's pairing. For example, devices that use
  "Just Works" pairing have no protection against man-in-the-middle attacks.
- **Devices designed to block relays.** Some devices, such as digital car keys, measure
  distance or latency to detect relays. gattroam won't try to bypass these protections.

## Responsible use

Only use gattroam with devices you own or are authorized to use.
