---
title: <span class="badge object-type-class"></span> TraceqlFilter
---
# <span class="badge object-type-class"></span> TraceqlFilter

## Definition

```python
class TraceqlFilter:
    # Uniquely identify the filter, will not be used in the query generation
    id_val: str
    # The tag for the search filter, for example: .http.status_code, .service.name, status
    tag: typing.Optional[str]
    # The operator that connects the tag to the value, for example: =, >, !=, =~
    operator: typing.Optional[str]
    # The value for the search filter
    value: typing.Optional[typing.Union[str, list[str]]]
    # The type of the value, used for example to check whether we need to wrap the value in quotes when generating the query
    value_type: typing.Optional[str]
    # The scope of the filter, can either be unscoped/all scopes, resource or span
    scope: typing.Optional[tempo.TraceqlSearchScope]
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

 * <span class="badge builder"></span> [TraceqlFilter](./builder-TraceqlFilter.md)
