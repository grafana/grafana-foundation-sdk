---
title: <span class="badge builder"></span> AdhocVariableBuilder
---
# <span class="badge builder"></span> AdhocVariableBuilder

## Constructor

```java
new AdhocVariableBuilder(String name)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public AdhocVariable build()
```

### <span class="badge object-method"></span> allowCustomValue

```java
public AdhocVariableBuilder allowCustomValue(Boolean allowCustomValue)
```

### <span class="badge object-method"></span> baseFilters

```java
public AdhocVariableBuilder baseFilters(List<com.grafana.foundation.cog.Builder<AdHocFilterWithLabels>> baseFilters)
```

### <span class="badge object-method"></span> datasource

```java
public AdhocVariableBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1AdhocVariableKindDatasource> datasource)
```

### <span class="badge object-method"></span> defaultKeys

```java
public AdhocVariableBuilder defaultKeys(List<com.grafana.foundation.cog.Builder<MetricFindValue>> defaultKeys)
```

### <span class="badge object-method"></span> description

```java
public AdhocVariableBuilder description(String description)
```

### <span class="badge object-method"></span> filters

```java
public AdhocVariableBuilder filters(List<com.grafana.foundation.cog.Builder<AdHocFilterWithLabels>> filters)
```

### <span class="badge object-method"></span> group

```java
public AdhocVariableBuilder group(String group)
```

### <span class="badge object-method"></span> hide

```java
public AdhocVariableBuilder hide(VariableHide hide)
```

### <span class="badge object-method"></span> label

```java
public AdhocVariableBuilder label(String label)
```

### <span class="badge object-method"></span> name

```java
public AdhocVariableBuilder name(String name)
```

### <span class="badge object-method"></span> skipUrlSync

```java
public AdhocVariableBuilder skipUrlSync(Boolean skipUrlSync)
```

## See also

 * <span class="badge object-type-class"></span> [AdhocVariableKind](./object-AdhocVariableKind.md)
