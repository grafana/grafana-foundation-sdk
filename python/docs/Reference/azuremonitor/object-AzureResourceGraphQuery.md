---
title: <span class="badge object-type-class"></span> AzureResourceGraphQuery
---
# <span class="badge object-type-class"></span> AzureResourceGraphQuery

## Definition

```python
class AzureResourceGraphQuery:
    # Azure Resource Graph KQL query to be executed.
    query: typing.Optional[str]
    # Specifies the format results should be returned as. Defaults to table.
    result_format: typing.Optional[str]
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

 * <span class="badge builder"></span> [AzureResourceGraphQuery](./builder-AzureResourceGraphQuery.md)
