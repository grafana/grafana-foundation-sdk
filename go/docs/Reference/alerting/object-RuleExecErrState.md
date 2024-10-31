---
title: <span class="badge object-type-enum"></span> RuleExecErrState
---
# <span class="badge object-type-enum"></span> RuleExecErrState

## Definition

```go
type RuleExecErrState string
const (
	RuleExecErrStateOK RuleExecErrState = "OK"
	RuleExecErrStateAlerting RuleExecErrState = "Alerting"
	RuleExecErrStateError RuleExecErrState = "Error"
)

```
