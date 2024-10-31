---
title: <span class="badge object-type-class"></span> ExprTypeResampleResultAssertions
---
# <span class="badge object-type-class"></span> ExprTypeResampleResultAssertions

## Definition

```python
class ExprTypeResampleResultAssertions:
    # Maximum frame count
    max_frames: typing.Optional[int]
    # Type asserts that the frame matches a known type structure.
    # Possible enum values:
    #  - `""` 
    #  - `"timeseries-wide"` 
    #  - `"timeseries-long"` 
    #  - `"timeseries-many"` 
    #  - `"timeseries-multi"` 
    #  - `"directory-listing"` 
    #  - `"table"` 
    #  - `"numeric-wide"` 
    #  - `"numeric-multi"` 
    #  - `"numeric-long"` 
    #  - `"log-lines"` 
    type_val: typing.Optional[typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]]
    # TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
    # contract documentation https://grafana.github.io/dataplane/contract/.
    type_version: list[int]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> [ExprTypeResampleResultAssertions](./builder-ExprTypeResampleResultAssertions.md)
