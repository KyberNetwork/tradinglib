# MEV Integration Notes

This file documents the Bombora MEV bundle sender, added in PR #218.
Read [CLAUDE.md](../CLAUDE.md) first for the general `pkg/mev` architecture.

## Bombora as a builder

Bombora is a new entry in the `BundleSenderType` enum in `pkg/mev/pkg.go`.
It follows the three-list pattern that every builder must follow:

1. `BundleSenderTypeBombora` in the `BundleSenderType` enum (`pkg/mev/pkg.go`).
2. `EndpointBombora = "https://rpc.bombora.build"` in `pkg/mev/broadcaster.go`.
3. `BuilderBomboraID = "builder-bombora"` in `pkg/mev/broadcaster.go`.

`pkg/mev/bundlesendertype_enumer.go` is generated.
The PR regenerated it after the enum change.

## Bombora-only bundle fields

`SendBundleV2Request` in `pkg/mev/pkg.go` gained 5 Bombora-only fields:

| Field | Type | Wire key | Meaning |
| --- | --- | --- | --- |
| `DroppingTxs` | `*[]string` | `droppingTxHashes` | Tx hashes that the builder can omit from the bundle, but must never include in a reverted state. |
| `ReplacementSeqNumber` | `*uint64` | `replacementSeqNumber` | Monotonically increasing sequence number for bundles that share one replacement UUID. |
| `RefundPercent` | `*uint64` | `refundPercent` | Percent of the bundle MEV profit to refund, 0 to 99. |
| `RefundRecipient` | `string` | `refundRecipient` | Address that receives the refund. |
| `RefundTxHashes` | `*[]string` | `refundTxHashes` | At most one tx hash that anchors the refund calculation. |

`SendBundleParams` (the wire type) carries the same 5 fields under a "Bombora-only bundle fields" comment block, placed after the pre-existing `StateBlockNumber` field.

## How the fields reach the wire

`Client.SendBundleV2` in `pkg/mev/bundle_sender.go` calls `SendBundleParams.SetBomboraFields(req)` only when `s.senderType == BundleSenderTypeBombora`.
Other builders never see these 5 fields on the wire, because `SetBomboraFields` runs only for that one sender type.

`SetBomboraFields` does 2 checks and appends an error to `p.Errors` on failure, instead of returning an error directly:

- If `RefundPercent` is above `BomboraMaxRefundPercent` (99, in `pkg/mev/constants.go`), it appends `ErrInvalidRefundPercent`.
- If `RefundTxHashes` holds more than 1 hash, it appends `ErrInvalidLenRefundTxHashes`.

Both errors live in `pkg/mev/errors.go`.
`SendBundleV2` calls `p.Err()` right after building `p`, so a validation failure surfaces before the HTTP call, not after.

## The `ReplacementUuid` wire key fix

`SendBundleParams.ReplacementUUID` changed its JSON tag from `ReplacementUuid` to `replacementUuid` (lowercase first letter).
This matches the real wire key that `CancelBundleParams` already used, and that Bombora's spec requires.

This tag is shared code: every sender that reaches `SetUUID` (`pkg/mev/bundle_sender.go`) and is not Beaver, Loki, or Jetbldr writes to `ReplacementUUID`, so this fix changes the wire key those senders send, not only Bombora's.
Beaver, Loki, and Jetbldr write to the separate `UUID` field (wire key `uuid`) instead, so they are unaffected.

## Adding another builder with refund or replacement fields

If a future builder needs fields similar to Bombora's, follow the same pattern:

1. Add request fields to `SendBundleV2Request`, each documented with a `// ... Builder-only.` comment.
2. Add matching wire fields to `SendBundleParams`.
3. Write a `SetXFields(req SendBundleV2Request) *SendBundleParams` method that copies and validates the fields, appending to `p.Errors` on a validation failure.
4. Call that method from `SendBundleV2`, gated on `s.senderType == BundleSenderTypeX`, so other builders never see the fields on the wire.
