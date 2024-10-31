---
title: <span class="badge object-type-class"></span> MatcherConfig
---
# <span class="badge object-type-class"></span> MatcherConfig

Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.

It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.

## Definition

```python
class MatcherConfig:
    """
    Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.
    It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.
    """

    # The matcher id. This is used to find the matcher implementation from registry.
    id_val: str
    # The matcher options. This is specific to the matcher implementation.
    options: typing.Optional[object]
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

