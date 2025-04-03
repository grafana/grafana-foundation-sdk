---
title: <span class="badge object-type-class"></span> TableFooterOptions
---
# <span class="badge object-type-class"></span> TableFooterOptions

Footer options

## Definition

```python
class TableFooterOptions:
    """
    Footer options
    """

    show: bool
    # actually 1 value
    reducer: list[str]
    fields: typing.Optional[list[str]]
    enable_pagination: typing.Optional[bool]
    count_rows: typing.Optional[bool]
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

 * <span class="badge builder"></span> [TableFooterOptions](./builder-TableFooterOptions.md)
