---
title: <span class="badge object-type-class"></span> DataTransformerConfig
---
# <span class="badge object-type-class"></span> DataTransformerConfig

Transformations allow to manipulate data returned by a query before the system applies a visualization.

Using transformations you can: rename fields, join time series data, perform mathematical operations across queries,

use the output of one transformation as the input to another transformation, etc.

## Definition

```python
class DataTransformerConfig:
    """
    Transformations allow to manipulate data returned by a query before the system applies a visualization.
    Using transformations you can: rename fields, join time series data, perform mathematical operations across queries,
    use the output of one transformation as the input to another transformation, etc.
    """

    # Unique identifier of transformer
    id_val: str
    # Disabled transformations are skipped
    disabled: typing.Optional[bool]
    # Optional frame matcher. When missing it will be applied to all results
    filter_val: typing.Optional[dashboard.MatcherConfig]
    # Options to be passed to the transformer
    # Valid options depend on the transformer id
    options: object
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

