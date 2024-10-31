---
title: <span class="badge object-type-class"></span> FieldConfigSource
---
# <span class="badge object-type-class"></span> FieldConfigSource

The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.

Each column within this structure is called a field. A field can represent a single time series or table column.

Field options allow you to change how the data is displayed in your visualizations.

## Definition

```python
class FieldConfigSource:
    """
    The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.
    Each column within this structure is called a field. A field can represent a single time series or table column.
    Field options allow you to change how the data is displayed in your visualizations.
    """

    # Defaults are the options applied to all fields.
    defaults: dashboard.FieldConfig
    # Overrides are the options applied to specific fields overriding the defaults.
    overrides: list[dashboard.DashboardFieldConfigSourceOverrides]
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

