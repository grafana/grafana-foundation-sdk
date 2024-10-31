---
title: <span class="badge object-type-class"></span> ElasticsearchTermsSettings
---
# <span class="badge object-type-class"></span> ElasticsearchTermsSettings

## Definition

```python
class ElasticsearchTermsSettings:
    order: typing.Optional[elasticsearch.TermsOrder]
    size: typing.Optional[str]
    min_doc_count: typing.Optional[str]
    order_by: typing.Optional[str]
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

 * <span class="badge builder"></span> [ElasticsearchTermsSettings](./builder-ElasticsearchTermsSettings.md)
