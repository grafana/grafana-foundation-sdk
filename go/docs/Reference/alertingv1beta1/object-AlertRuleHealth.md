---
title: <span class="badge object-type-enum"></span> AlertRuleHealth
---
# <span class="badge object-type-enum"></span> AlertRuleHealth

## Definition

```go
type AlertRuleHealth string
const (
	AlertRuleHealthUnknown AlertRuleHealth = "Unknown"
	AlertRuleHealthOK AlertRuleHealth = "OK"
	AlertRuleHealthPaused AlertRuleHealth = "Paused"
	AlertRuleHealthError AlertRuleHealth = "Error"
	AlertRuleHealthNoData AlertRuleHealth = "NoData"
)

```
