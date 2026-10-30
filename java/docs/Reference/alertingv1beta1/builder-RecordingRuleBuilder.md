---
title: <span class="badge builder"></span> RecordingRuleBuilder
---
# <span class="badge builder"></span> RecordingRuleBuilder

## Constructor

```java
new RecordingRuleBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public RecordingRule build()
```

### <span class="badge object-method"></span> expressions

```java
public RecordingRuleBuilder expressions(Map<String, com.grafana.foundation.cog.Builder<Expression>> expressions)
```

### <span class="badge object-method"></span> labels

```java
public RecordingRuleBuilder labels(Map<String, String> labels)
```

### <span class="badge object-method"></span> metric

```java
public RecordingRuleBuilder metric(String metric)
```

### <span class="badge object-method"></span> paused

```java
public RecordingRuleBuilder paused(Boolean paused)
```

### <span class="badge object-method"></span> targetDatasourceUID

```java
public RecordingRuleBuilder targetDatasourceUID(String targetDatasourceUID)
```

### <span class="badge object-method"></span> title

```java
public RecordingRuleBuilder title(String title)
```

### <span class="badge object-method"></span> trigger

```java
public RecordingRuleBuilder trigger(com.grafana.foundation.cog.Builder<IntervalTrigger> trigger)
```

## See also

 * <span class="badge object-type-class"></span> [RecordingRuleSpec](./object-RecordingRuleSpec.md)
