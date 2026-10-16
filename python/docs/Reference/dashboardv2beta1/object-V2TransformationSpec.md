---
title: <span class="badge object-type-class"></span> V2TransformationSpec
---
# <span class="badge object-type-class"></span> V2TransformationSpec

Dashboard v2 shape: no `id`, it is carried by the parent's `group`.

## Definition

```python
class V2TransformationSpec:
    """
    Dashboard v2 shape: no `id`, it is carried by the parent's `group`.
    """

    # Disabled transformations are skipped
    disabled: typing.Optional[bool]
    # Optional frame matcher. When missing it will be applied to all results
    filter_val: typing.Optional[dashboardv2beta1.MatcherConfig]
    # Where to pull DataFrames from as input to transformation
    topic: typing.Optional[dashboardv2beta1.DataTopic]
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

## See also

 * <span class="badge builder"></span> [V2TransformationSpec](./builder-V2TransformationSpec.md)
