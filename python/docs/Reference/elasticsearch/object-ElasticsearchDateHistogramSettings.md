---
title: <span class="badge object-type-class"></span> ElasticsearchDateHistogramSettings
---
# <span class="badge object-type-class"></span> ElasticsearchDateHistogramSettings

## Definition

```python
class ElasticsearchDateHistogramSettings:
    interval: typing.Optional[str]
    min_doc_count: typing.Optional[str]
    trim_edges: typing.Optional[str]
    offset: typing.Optional[str]
    time_zone: typing.Optional[str]
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

 * <span class="badge builder"></span> [ElasticsearchDateHistogramSettings](./builder-ElasticsearchDateHistogramSettings.md)
