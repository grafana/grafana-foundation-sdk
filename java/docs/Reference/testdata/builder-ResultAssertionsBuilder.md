---
title: <span class="badge builder"></span> ResultAssertionsBuilder
---
# <span class="badge builder"></span> ResultAssertionsBuilder

## Constructor

```java
new ResultAssertionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public ResultAssertions build()
```

### <span class="badge object-method"></span> maxFrames

Maximum frame count

```java
public ResultAssertionsBuilder maxFrames(Long maxFrames)
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
public ResultAssertionsBuilder type(ResultAssertionsType type)
```

### <span class="badge object-method"></span> typeVersion

TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane

contract documentation https://grafana.github.io/dataplane/contract/.

```java
public ResultAssertionsBuilder typeVersion(List<Long> typeVersion)
```

## See also

 * <span class="badge object-type-class"></span> [ResultAssertions](./object-ResultAssertions.md)
