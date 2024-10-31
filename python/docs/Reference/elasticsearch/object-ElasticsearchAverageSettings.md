---
title: <span class="badge object-type-class"></span> ElasticsearchAverageSettings
---
# <span class="badge object-type-class"></span> ElasticsearchAverageSettings

## Definition

```python
class ElasticsearchAverageSettings:
    script: typing.Optional[elasticsearch.InlineScript]
    missing: typing.Optional[str]
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

 * <span class="badge builder"></span> [ElasticsearchAverageSettings](./builder-ElasticsearchAverageSettings.md)
