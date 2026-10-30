---
title: <span class="badge builder"></span> AlertRuleStatusBuilder
---
# <span class="badge builder"></span> AlertRuleStatusBuilder

## Constructor

```java
new AlertRuleStatusBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public AlertRuleStatus build()
```

### <span class="badge object-method"></span> evaluationDuration

duration of the last evaluation in seconds

```java
public AlertRuleStatusBuilder evaluationDuration(Double evaluationDuration)
```

### <span class="badge object-method"></span> health

```java
public AlertRuleStatusBuilder health(AlertRuleHealth health)
```

### <span class="badge object-method"></span> lastError

```java
public AlertRuleStatusBuilder lastError(String lastError)
```

### <span class="badge object-method"></span> lastEvaluationTime

```java
public AlertRuleStatusBuilder lastEvaluationTime(String lastEvaluationTime)
```

### <span class="badge object-method"></span> state

```java
public AlertRuleStatusBuilder state(AlertRuleState state)
```

### <span class="badge object-method"></span> stateReason

```java
public AlertRuleStatusBuilder stateReason(AlertRuleStateReason stateReason)
```

### <span class="badge object-method"></span> totals

```java
public AlertRuleStatusBuilder totals(com.grafana.foundation.cog.Builder<AlertRuleInstanceTotals> totals)
```

## See also

 * <span class="badge object-type-class"></span> [AlertRuleStatus](./object-AlertRuleStatus.md)
