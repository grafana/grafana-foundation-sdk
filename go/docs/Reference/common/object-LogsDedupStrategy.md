---
title: <span class="badge object-type-enum"></span> LogsDedupStrategy
---
# <span class="badge object-type-enum"></span> LogsDedupStrategy

## Definition

```go
type LogsDedupStrategy string
const (
	LogsDedupStrategyNone LogsDedupStrategy = "none"
	LogsDedupStrategyExact LogsDedupStrategy = "exact"
	LogsDedupStrategyNumbers LogsDedupStrategy = "numbers"
	LogsDedupStrategySignature LogsDedupStrategy = "signature"
)

```
