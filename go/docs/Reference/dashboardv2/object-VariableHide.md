---
title: <span class="badge object-type-enum"></span> VariableHide
---
# <span class="badge object-type-enum"></span> VariableHide

Determine if the variable shows on dashboard

Accepted values are `dontHide` (show label and value), `hideLabel` (show value only), `hideVariable` (show nothing), `inControlsMenu` (show in a drop-down menu).

## Definition

```go
type VariableHide string
const (
	VariableHideDontHide VariableHide = "dontHide"
	VariableHideHideLabel VariableHide = "hideLabel"
	VariableHideHideVariable VariableHide = "hideVariable"
	VariableHideInControlsMenu VariableHide = "inControlsMenu"
)

```
