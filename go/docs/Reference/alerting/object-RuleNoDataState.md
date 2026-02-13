---
title: <span class="badge object-type-enum"></span> RuleNoDataState
---
# <span class="badge object-type-enum"></span> RuleNoDataState

## Definition

```go
type RuleNoDataState string
const (
	RuleNoDataStateOK RuleNoDataState = "OK"
	RuleNoDataStateAlerting RuleNoDataState = "Alerting"
	RuleNoDataStateNoData RuleNoDataState = "NoData"
	RuleNoDataStateKeepLast RuleNoDataState = "KeepLast"
)

```
