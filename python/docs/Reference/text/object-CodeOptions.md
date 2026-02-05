---
title: <span class="badge object-type-class"></span> CodeOptions
---
# <span class="badge object-type-class"></span> CodeOptions

## Definition

```python
class CodeOptions:
    # The language passed to monaco code editor
    language: text.CodeLanguage
    show_line_numbers: bool
    show_mini_map: bool
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

 * <span class="badge builder"></span> [CodeOptions](./builder-CodeOptions.md)
