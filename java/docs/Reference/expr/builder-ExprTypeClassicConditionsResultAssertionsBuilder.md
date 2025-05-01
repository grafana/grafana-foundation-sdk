---
title: <span class="badge builder"></span> ExprTypeClassicConditionsResultAssertionsBuilder
---
# <span class="badge builder"></span> ExprTypeClassicConditionsResultAssertionsBuilder

## Constructor

```java
new ExprTypeClassicConditionsResultAssertionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public ExprTypeClassicConditionsResultAssertions build()
```

### <span class="badge object-method"></span> maxFrames

Maximum frame count

```java
public ExprTypeClassicConditionsResultAssertionsBuilder maxFrames(Long maxFrames)
```

### <span class="badge object-method"></span> type

Type asserts that the frame matches a known type structure.

Possible enum values:

 - `""` 

 - `"timeseries-wide"` 

 - `"timeseries-long"` 

 - `"timeseries-many"` 

 - `"timeseries-multi"` 

 - `"directory-listing"` 

 - `"table"` 

 - `"numeric-wide"` 

 - `"numeric-multi"` 

 - `"numeric-long"` 

 - `"log-lines"` 

```java
public ExprTypeClassicConditionsResultAssertionsBuilder type(ExprTypeClassicConditionsResultAssertionsType type)
```

### <span class="badge object-method"></span> typeVersion

TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane

contract documentation https://grafana.github.io/dataplane/contract/.

```java
public ExprTypeClassicConditionsResultAssertionsBuilder typeVersion(List<Long> typeVersion)
```

## See also

 * <span class="badge object-type-class"></span> [ExprTypeClassicConditionsResultAssertions](./object-ExprTypeClassicConditionsResultAssertions.md)
