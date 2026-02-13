---
title: <span class="badge object-type-enum"></span> RuleExecErrState
---
# <span class="badge object-type-enum"></span> RuleExecErrState

## Definition

```go
type RuleExecErrState string
const (
	RuleExecErrStateAlerting RuleExecErrState = "Alerting"
	RuleExecErrStateError RuleExecErrState = "Error"
	RuleExecErrStateOK RuleExecErrState = "OK"
	RuleExecErrStateKeepLast RuleExecErrState = "KeepLast"
)

```
