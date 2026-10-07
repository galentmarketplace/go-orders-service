# go-orders-service

A small Go service used to demonstrate the AI Testing Marketplace's coverage agent.

`internal/pricing` has tests for `Subtotal` and `ApplyTax`, but **none** for
`BulkDiscount` or `LoyaltyPoints`. `internal/cart` has no tests at all.

The coverage agent measures this, writes tests for the uncovered functions,
verifies they compile and pass, re-measures, and opens a pull request.
