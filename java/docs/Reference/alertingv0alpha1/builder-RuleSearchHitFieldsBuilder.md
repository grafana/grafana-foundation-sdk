---
title: <span class="badge builder"></span> RuleSearchHitFieldsBuilder
---
# <span class="badge builder"></span> RuleSearchHitFieldsBuilder

#RuleSearchHitFields is the per-kind field payload returned on each hit.

It carries the union of alert- and recording-rule search fields; only the

fields relevant to a hit's kind are populated. This maps to the kind's

declared searchFields.

## Constructor

```java
new RuleSearchHitFieldsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public RuleSearchHitFields build()
```

### <span class="badge object-method"></span> annotations

Alert-rule fields.

```java
public RuleSearchHitFieldsBuilder annotations(Map<String, String> annotations)
```

### <span class="badge object-method"></span> dashboardUID

```java
public RuleSearchHitFieldsBuilder dashboardUID(String dashboardUID)
```

### <span class="badge object-method"></span> datasourceUIDs

```java
public RuleSearchHitFieldsBuilder datasourceUIDs(List<String> datasourceUIDs)
```

### <span class="badge object-method"></span> folder

```java
public RuleSearchHitFieldsBuilder folder(String folder)
```

### <span class="badge object-method"></span> for

```java
public RuleSearchHitFieldsBuilder for(String forArg)
```

### <span class="badge object-method"></span> interval

```java
public RuleSearchHitFieldsBuilder interval(String interval)
```

### <span class="badge object-method"></span> keepFiringFor

```java
public RuleSearchHitFieldsBuilder keepFiringFor(String keepFiringFor)
```

### <span class="badge object-method"></span> labels

```java
public RuleSearchHitFieldsBuilder labels(Map<String, String> labels)
```

### <span class="badge object-method"></span> metric

Recording-rule fields.

```java
public RuleSearchHitFieldsBuilder metric(String metric)
```

### <span class="badge object-method"></span> notificationType

```java
public RuleSearchHitFieldsBuilder notificationType(String notificationType)
```

### <span class="badge object-method"></span> panelID

```java
public RuleSearchHitFieldsBuilder panelID(Long panelID)
```

### <span class="badge object-method"></span> paused

```java
public RuleSearchHitFieldsBuilder paused(Boolean paused)
```

### <span class="badge object-method"></span> receiver

```java
public RuleSearchHitFieldsBuilder receiver(String receiver)
```

### <span class="badge object-method"></span> routingTree

```java
public RuleSearchHitFieldsBuilder routingTree(String routingTree)
```

### <span class="badge object-method"></span> targetDatasourceUID

```java
public RuleSearchHitFieldsBuilder targetDatasourceUID(String targetDatasourceUID)
```

### <span class="badge object-method"></span> title

```java
public RuleSearchHitFieldsBuilder title(String title)
```

### <span class="badge object-method"></span> type

```java
public RuleSearchHitFieldsBuilder type(String type)
```

## See also

 * <span class="badge object-type-class"></span> [RuleSearchHitFields](./object-RuleSearchHitFields.md)
