---
title: <span class="badge builder"></span> RuleBuilder
---
# <span class="badge builder"></span> RuleBuilder

## Constructor

```java
new RuleBuilder(String title)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Rule build()
```

### <span class="badge object-method"></span> annotations

```java
public RuleBuilder annotations(Map<String, String> annotations)
```

### <span class="badge object-method"></span> condition

```java
public RuleBuilder condition(String condition)
```

### <span class="badge object-method"></span> execErrState

```java
public RuleBuilder execErrState(RuleExecErrState execErrState)
```

### <span class="badge object-method"></span> folderUID

```java
public RuleBuilder folderUID(String folderUID)
```

### <span class="badge object-method"></span> for

The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.

Before this time has elapsed, the rule is only considered to be Pending.

```java
public RuleBuilder for(String forArg)
```

### <span class="badge object-method"></span> id

```java
public RuleBuilder id(Long id)
```

### <span class="badge object-method"></span> isPaused

```java
public RuleBuilder isPaused(Boolean isPaused)
```

### <span class="badge object-method"></span> labels

```java
public RuleBuilder labels(Map<String, String> labels)
```

### <span class="badge object-method"></span> noDataState

```java
public RuleBuilder noDataState(RuleNoDataState noDataState)
```

### <span class="badge object-method"></span> notificationSettings

```java
public RuleBuilder notificationSettings(com.grafana.foundation.cog.Builder<NotificationSettings> notificationSettings)
```

### <span class="badge object-method"></span> orgID

```java
public RuleBuilder orgID(Long orgID)
```

### <span class="badge object-method"></span> provenance

```java
public RuleBuilder provenance(String provenance)
```

### <span class="badge object-method"></span> queries

```java
public RuleBuilder queries(List<com.grafana.foundation.cog.Builder<Query>> data)
```

### <span class="badge object-method"></span> ruleGroup

```java
public RuleBuilder ruleGroup(String ruleGroup)
```

### <span class="badge object-method"></span> title

```java
public RuleBuilder title(String title)
```

### <span class="badge object-method"></span> uid

```java
public RuleBuilder uid(String uid)
```

### <span class="badge object-method"></span> updated

```java
public RuleBuilder updated(String updated)
```

### <span class="badge object-method"></span> withQuery

```java
public RuleBuilder withQuery(com.grafana.foundation.cog.Builder<Query> data)
```

## See also

 * <span class="badge object-type-class"></span> [Rule](./object-Rule.md)
