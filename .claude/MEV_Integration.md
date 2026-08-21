# MEV builder integration guide

This file is a guide for adding a new builder to `pkg/mev`.
Read [CLAUDE.md](../CLAUDE.md) first for the general `pkg/mev` architecture.
The Bombora builder (PR #218) is used below as a worked example only.
Do not copy Bombora-specific names for a different builder.

## Step 1: register the builder in the 3 parallel lists

Every builder needs an entry in each of these 3 lists, and the lists must stay in sync by hand:

1. A `BundleSenderType` value in the enum in `pkg/mev/pkg.go`.
2. An `Endpoint*` constant in `pkg/mev/broadcaster.go`, for example `EndpointBombora = "https://rpc.bombora.build"`.
3. A `Builder*ID` constant in `pkg/mev/broadcaster.go`, for example `BuilderBomboraID = "builder-bombora"`.

After you change the enum, regenerate `pkg/mev/bundlesendertype_enumer.go`:

1. Run `go install github.com/dmarkham/enumer@latest`.
2. Run `enumer -type=BundleSenderType -linecomment` from `pkg/mev`.

Do not edit the generated file by hand.

## Step 2: decide whether the builder needs its own client file

Most builders reuse the generic `Client` in `pkg/mev/bundle_sender.go`.
Add a separate client file, such as `blxr_bundle_sender.go` or `merkle_sender.go`, only when the builder's request or auth shape does not fit `Client`.

## Step 3: add builder-only bundle fields, if the builder needs them

Some builders accept extra fields on `eth_sendBundle` that other builders do not support.
Bombora is the current example: it added `DroppingTxs`, `ReplacementSeqNumber`, `RefundPercent`, `RefundRecipient`, and `RefundTxHashes`.

Follow this pattern for a new builder with similar needs:

1. Add the new fields to `SendBundleV2Request` in `pkg/mev/pkg.go`.
   Comment each field with its purpose and the builder name, for example `// ... RefundPercent-style field. Bombora only.`
2. Add matching wire fields to `SendBundleParams` in `pkg/mev/bundle_sender.go`, grouped under a comment naming the builder, for example `// Bombora-only bundle fields.`
   Place the new group after existing fields; do not reorder pre-existing fields, since that makes the diff imply they are new.
3. Write one `SetXFields(req SendBundleV2Request) *SendBundleParams` method on `SendBundleParams` that copies the fields from `req` and validates them.
   On a validation failure, append a sentinel error to `p.Errors` instead of returning an error directly, so the method keeps returning `*SendBundleParams` for chaining.
4. Add any new sentinel errors to `pkg/mev/errors.go`.
5. Add any new numeric limits, such as a maximum percent, as named constants in `pkg/mev/constants.go`.
6. In `Client.SendBundleV2`, call the new `SetXFields` method gated on `s.senderType == BundleSenderTypeX`, so other builders never see the new fields on the wire.

`SendBundleV2` calls `p.Err()` right after building `p`, so a validation failure appended to `p.Errors` surfaces before the HTTP call.

## Step 4: check shared code before changing a JSON tag

Some fields on `SendBundleParams`, such as `ReplacementUUID`, are shared across builders through `SetUUID`.
Before you change a JSON tag or a field's meaning, check every sender type that reaches the same field, not only the builder you are adding.
For example, `SetUUID` writes to a different field (`UUID`, wire key `uuid`) for Beaver, Loki, and Jetbldr, and to `ReplacementUUID` for every other sender type; a tag change on `ReplacementUUID` therefore changes the wire output for all of those other senders at once.

## Step 5: verify

Run, from the repository root:

```sh
go build ./...
go test -race ./...
golangci-lint run --config=.golangci.yml
```

Do not add a new automated test file for the builder's HTTP calls.
Existing builder clients use a manual `t.Skip()` test that needs a live builder endpoint; follow that pattern instead of an `httptest` harness.
