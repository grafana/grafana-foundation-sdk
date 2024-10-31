---
title: <span class="badge object-type-enum"></span> VariableHide
---
# <span class="badge object-type-enum"></span> VariableHide

Determine if the variable shows on dashboard

Accepted values are 0 (show label and value), 1 (show value only), 2 (show nothing).

## Definition

```go
type VariableHide int64
const (
	VariableHideDontHide VariableHide = 0
	VariableHideHideLabel VariableHide = 1
	VariableHideHideVariable VariableHide = 2
)

```
