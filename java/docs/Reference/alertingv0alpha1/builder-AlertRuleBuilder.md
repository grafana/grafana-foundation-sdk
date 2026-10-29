---
title: <span class="badge builder"></span> AlertRuleBuilder
---
# <span class="badge builder"></span> AlertRuleBuilder

## Constructor

```java
new AlertRuleBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public AlertRule build()
```

### <span class="badge object-method"></span> annotations

```java
public AlertRuleBuilder annotations(Map<String, String> annotations)
```

### <span class="badge object-method"></span> execErrState

```java
public AlertRuleBuilder execErrState(ExecErrState execErrState)
```

### <span class="badge object-method"></span> expressions

```java
public AlertRuleBuilder expressions(Map<String, com.grafana.foundation.cog.Builder<Expression>> expressions)
```

### <span class="badge object-method"></span> for

```java
public AlertRuleBuilder for(String forArg)
```

### <span class="badge object-method"></span> keepFiringFor

```java
public AlertRuleBuilder keepFiringFor(String keepFiringFor)
```

### <span class="badge object-method"></span> labels

```java
public AlertRuleBuilder labels(Map<String, String> labels)
```

### <span class="badge object-method"></span> missingSeriesEvalsToResolve

```java
public AlertRuleBuilder missingSeriesEvalsToResolve(Long missingSeriesEvalsToResolve)
```

### <span class="badge object-method"></span> noDataState

```java
public AlertRuleBuilder noDataState(NoDataState noDataState)
```

### <span class="badge object-method"></span> notificationSettings

```java
public AlertRuleBuilder notificationSettings(NotificationSettings notificationSettings)
```

### <span class="badge object-method"></span> panelRef

```java
public AlertRuleBuilder panelRef(com.grafana.foundation.cog.Builder<PanelRef> panelRef)
```

### <span class="badge object-method"></span> paused

```java
public AlertRuleBuilder paused(Boolean paused)
```

### <span class="badge object-method"></span> title

```java
public AlertRuleBuilder title(String title)
```

### <span class="badge object-method"></span> trigger

```java
public AlertRuleBuilder trigger(com.grafana.foundation.cog.Builder<IntervalTrigger> trigger)
```

## See also

 * <span class="badge object-type-class"></span> [AlertRuleSpec](./object-AlertRuleSpec.md)
