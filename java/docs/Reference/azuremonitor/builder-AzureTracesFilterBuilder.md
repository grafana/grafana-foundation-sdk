---
title: <span class="badge builder"></span> AzureTracesFilterBuilder
---
# <span class="badge builder"></span> AzureTracesFilterBuilder

## Constructor

```java
new AzureTracesFilterBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public AzureTracesFilter build()
```

### <span class="badge object-method"></span> filters

Values to filter by.

```java
public AzureTracesFilterBuilder filters(List<String> filters)
```

### <span class="badge object-method"></span> operation

Comparison operator to use. Either equals or not equals.

```java
public AzureTracesFilterBuilder operation(String operation)
```

### <span class="badge object-method"></span> property

Property name, auto-populated based on available traces.

```java
public AzureTracesFilterBuilder property(String property)
```

## See also

 * <span class="badge object-type-class"></span> [AzureTracesFilter](./object-AzureTracesFilter.md)
