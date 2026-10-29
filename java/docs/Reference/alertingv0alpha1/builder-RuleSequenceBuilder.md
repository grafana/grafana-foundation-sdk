---
title: <span class="badge builder"></span> RuleSequenceBuilder
---
# <span class="badge builder"></span> RuleSequenceBuilder

## Constructor

```java
new RuleSequenceBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public RuleSequence build()
```

### <span class="badge object-method"></span> alertingRules

```java
public RuleSequenceBuilder alertingRules(List<com.grafana.foundation.cog.Builder<RuleRef>> alertingRules)
```

### <span class="badge object-method"></span> recordingRules

```java
public RuleSequenceBuilder recordingRules(List<com.grafana.foundation.cog.Builder<RuleRef>> recordingRules)
```

### <span class="badge object-method"></span> trigger

```java
public RuleSequenceBuilder trigger(com.grafana.foundation.cog.Builder<IntervalTrigger> trigger)
```

## See also

 * <span class="badge object-type-class"></span> [RuleSequenceSpec](./object-RuleSequenceSpec.md)
