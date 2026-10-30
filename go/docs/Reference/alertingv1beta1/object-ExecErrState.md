---
title: <span class="badge object-type-enum"></span> ExecErrState
---
# <span class="badge object-type-enum"></span> ExecErrState

## Definition

```go
type ExecErrState string
const (
	ExecErrStateError ExecErrState = "Error"
	ExecErrStateOk ExecErrState = "Ok"
	ExecErrStateAlerting ExecErrState = "Alerting"
	ExecErrStateKeepLast ExecErrState = "KeepLast"
)

```
