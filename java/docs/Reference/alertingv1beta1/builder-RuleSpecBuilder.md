---
title: <span class="badge builder"></span> RuleSpecBuilder
---
# <span class="badge builder"></span> RuleSpecBuilder

## Constructor

```java
new RuleSpecBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public RuleSpec build()
```

### <span class="badge object-method"></span> expressions

```java
public RuleSpecBuilder expressions(Map<String, com.grafana.foundation.cog.Builder<Expression>> expressions)
```

### <span class="badge object-method"></span> labels

```java
public RuleSpecBuilder labels(Map<String, String> labels)
```

### <span class="badge object-method"></span> paused

```java
public RuleSpecBuilder paused(Boolean paused)
```

### <span class="badge object-method"></span> title

```java
public RuleSpecBuilder title(String title)
```

### <span class="badge object-method"></span> trigger

```java
public RuleSpecBuilder trigger(com.grafana.foundation.cog.Builder<IntervalTrigger> trigger)
```

## See also

 * <span class="badge object-type-class"></span> [RuleSpec](./object-RuleSpec.md)
