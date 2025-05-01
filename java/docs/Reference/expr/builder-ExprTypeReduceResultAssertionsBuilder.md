---
title: <span class="badge builder"></span> ExprTypeReduceResultAssertionsBuilder
---
# <span class="badge builder"></span> ExprTypeReduceResultAssertionsBuilder

## Constructor

```java
new ExprTypeReduceResultAssertionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public ExprTypeReduceResultAssertions build()
```

### <span class="badge object-method"></span> maxFrames

Maximum frame count

```java
public ExprTypeReduceResultAssertionsBuilder maxFrames(Long maxFrames)
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
public ExprTypeReduceResultAssertionsBuilder type(ExprTypeReduceResultAssertionsType type)
```

### <span class="badge object-method"></span> typeVersion

TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane

contract documentation https://grafana.github.io/dataplane/contract/.

```java
public ExprTypeReduceResultAssertionsBuilder typeVersion(List<Long> typeVersion)
```

## See also

 * <span class="badge object-type-class"></span> [ExprTypeReduceResultAssertions](./object-ExprTypeReduceResultAssertions.md)
