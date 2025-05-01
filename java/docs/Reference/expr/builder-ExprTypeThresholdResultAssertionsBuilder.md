---
title: <span class="badge builder"></span> ExprTypeThresholdResultAssertionsBuilder
---
# <span class="badge builder"></span> ExprTypeThresholdResultAssertionsBuilder

## Constructor

```java
new ExprTypeThresholdResultAssertionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public ExprTypeThresholdResultAssertions build()
```

### <span class="badge object-method"></span> maxFrames

Maximum frame count

```java
public ExprTypeThresholdResultAssertionsBuilder maxFrames(Long maxFrames)
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
public ExprTypeThresholdResultAssertionsBuilder type(ExprTypeThresholdResultAssertionsType type)
```

### <span class="badge object-method"></span> typeVersion

TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane

contract documentation https://grafana.github.io/dataplane/contract/.

```java
public ExprTypeThresholdResultAssertionsBuilder typeVersion(List<Long> typeVersion)
```

## See also

 * <span class="badge object-type-class"></span> [ExprTypeThresholdResultAssertions](./object-ExprTypeThresholdResultAssertions.md)
