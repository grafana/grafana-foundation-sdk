---
title: <span class="badge builder"></span> ResultAssertionsBuilder
---
# <span class="badge builder"></span> ResultAssertionsBuilder

## Constructor

```typescript
new ResultAssertionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> maxFrames

Maximum frame count

```typescript
maxFrames(maxFrames: number)
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

```typescript
type(type: "" | "timeseries-wide" | "timeseries-long" | "timeseries-many" | "timeseries-multi" | "directory-listing" | "table" | "numeric-wide" | "numeric-multi" | "numeric-long" | "log-lines")
```

### <span class="badge object-method"></span> typeVersion

TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane

contract documentation https://grafana.github.io/dataplane/contract/.

```typescript
typeVersion(typeVersion: number[])
```

## See also

 * <span class="badge object-type-interface"></span> [ResultAssertions](./object-ResultAssertions.md)
