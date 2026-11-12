# Architecture Decision Records

This folder records the important decisions behind gattroam: what was decided, why, and
what it costs. These records help new contributors understand the project, and keep
settled questions from being reopened without new information.

| # | Decision | Status |
|---|---|---|
| [0001](0001-start-with-ble.md) | Start with BLE (GATT) before Bluetooth Classic | Accepted |
| [0002](0002-use-bluez.md) | Use BlueZ, through its D-Bus API, as the Bluetooth stack | Accepted |
| [0003](0003-linux-only.md) | Support Linux only | Accepted |

## Writing a new ADR

1. Create a file named `NNNN-short-title.md`, using the next number and the template
   below.
2. Open a pull request. The discussion happens there.
3. Don't rewrite an accepted ADR. If a decision changes, write a new ADR and mark the old
   one as "Superseded by NNNN".

```markdown
# NNNN. Title

- **Status:** Proposed | Accepted | Superseded by NNNN
- **Date:** YYYY-MM-DD

## Context

What problem are we solving, and which constraints matter?

## Decision

What we decided to do.

## Consequences

What gets easier, what gets harder, and what we give up.
```
