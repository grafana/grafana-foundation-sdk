---
title: <span class="badge object-type-enum"></span> AlertRuleState
---
# <span class="badge object-type-enum"></span> AlertRuleState

## Definition

```go
type AlertRuleState string
const (
	AlertRuleStateInactive AlertRuleState = "Inactive"
	AlertRuleStateHealthy AlertRuleState = "Healthy"
	AlertRuleStateFiring AlertRuleState = "Firing"
	AlertRuleStatePending AlertRuleState = "Pending"
	AlertRuleStateRecovering AlertRuleState = "Recovering"
)

```
