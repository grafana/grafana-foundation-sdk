---
title: <span class="badge builder"></span> ExprTypeMathResultAssertions
---
# <span class="badge builder"></span> ExprTypeMathResultAssertions

## Constructor

```python
ExprTypeMathResultAssertions()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> expr.ExprTypeMathResultAssertions
```

### <span class="badge object-method"></span> max_frames

Maximum frame count

```python
def max_frames(max_frames: int) -> typing.Self
```

### <span class="badge object-method"></span> type_val

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

```python
def type_val(type_val: typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]) -> typing.Self
```

### <span class="badge object-method"></span> type_version

TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane

contract documentation https://grafana.github.io/dataplane/contract/.

```python
def type_version(type_version: list[int]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [ExprTypeMathResultAssertions](./object-ExprTypeMathResultAssertions.md)
