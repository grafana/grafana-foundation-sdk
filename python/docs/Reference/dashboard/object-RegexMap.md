---
title: <span class="badge object-type-class"></span> RegexMap
---
# <span class="badge object-type-class"></span> RegexMap

Maps regular expressions to replacement text and a color.

For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.

## Definition

```python
class RegexMap:
    """
    Maps regular expressions to replacement text and a color.
    For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
    """

    type_val: typing.Literal["regex"]
    # Regular expression to match against and the result to apply when the value matches the regex
    options: dashboard.DashboardRegexMapOptions
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

 * <span class="badge builder"></span> [RegexMap](./builder-RegexMap.md)
