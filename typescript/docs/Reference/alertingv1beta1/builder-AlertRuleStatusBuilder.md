---
title: <span class="badge builder"></span> AlertRuleStatusBuilder
---
# <span class="badge builder"></span> AlertRuleStatusBuilder

## Constructor

```typescript
new AlertRuleStatusBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> evaluationDuration

duration of the last evaluation in seconds

```typescript
evaluationDuration(evaluationDuration: number)
```

### <span class="badge object-method"></span> health

```typescript
health(health: alertingv1beta1.AlertRuleHealth)
```

### <span class="badge object-method"></span> lastError

```typescript
lastError(lastError: string)
```

### <span class="badge object-method"></span> lastEvaluationTime

```typescript
lastEvaluationTime(lastEvaluationTime: string)
```

### <span class="badge object-method"></span> state

```typescript
state(state: alertingv1beta1.AlertRuleState)
```

### <span class="badge object-method"></span> stateReason

```typescript
stateReason(stateReason: alertingv1beta1.AlertRuleStateReason)
```

### <span class="badge object-method"></span> totals

```typescript
totals(totals: cog.Builder<alertingv1beta1.AlertRuleInstanceTotals>)
```

## See also

 * <span class="badge object-type-interface"></span> [AlertRuleStatus](./object-AlertRuleStatus.md)
