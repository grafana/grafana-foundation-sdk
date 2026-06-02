---
title: <span class="badge builder"></span> TracesFilterBuilder
---
# <span class="badge builder"></span> TracesFilterBuilder

## Constructor

```java
new TracesFilterBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public TracesFilter build()
```

### <span class="badge object-method"></span> filters

Values to filter by.

```java
public TracesFilterBuilder filters(List<String> filters)
```

### <span class="badge object-method"></span> operation

Comparison operator to use. Either equals or not equals.

```java
public TracesFilterBuilder operation(String operation)
```

### <span class="badge object-method"></span> property

Property name, auto-populated based on available traces.

```java
public TracesFilterBuilder property(String property)
```

## See also

 * <span class="badge object-type-class"></span> [TracesFilter](./object-TracesFilter.md)
