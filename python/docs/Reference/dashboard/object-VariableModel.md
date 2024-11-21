---
title: <span class="badge object-type-class"></span> VariableModel
---
# <span class="badge object-type-class"></span> VariableModel

A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.

## Definition

```python
class VariableModel:
    """
    A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
    """

    # Unique numeric identifier for the variable.
    id_val: str
    # Type of variable
    type_val: dashboard.VariableType
    # Name of variable
    name: str
    # Optional display name
    label: typing.Optional[str]
    # Visibility configuration for the variable
    hide: dashboard.VariableHide
    # Whether the variable value should be managed by URL query params or not
    skip_url_sync: bool
    # Description of variable. It can be defined but `null`.
    description: typing.Optional[str]
    # Query used to fetch values for a variable
    query: typing.Optional[typing.Union[str, dict[str, object]]]
    # Data source used to fetch values for a variable. It can be defined but `null`.
    datasource: typing.Optional[dashboard.DataSourceRef]
    # Format to use while fetching all values from data source, eg: wildcard, glob, regex, pipe, etc.
    all_format: typing.Optional[str]
    # Shows current selected variable text/value on the dashboard
    current: typing.Optional[dashboard.VariableOption]
    # Whether multiple values can be selected or not from variable value list
    multi: typing.Optional[bool]
    # Options that can be selected for a variable.
    options: typing.Optional[list[dashboard.VariableOption]]
    # Options to config when to refresh a variable
    refresh: typing.Optional[dashboard.VariableRefresh]
    # Options sort order
    sort: typing.Optional[dashboard.VariableSort]
    # Whether all value option is available or not
    include_all: typing.Optional[bool]
    # Custom all value
    all_value: typing.Optional[str]
    # Optional field, if you want to extract part of a series name or metric node segment.
    # Named capture groups can be used to separate the display text and value.
    regex: typing.Optional[str]
    # Dynamically calculates interval by dividing time range by the count specified.
    auto: typing.Optional[bool]
    # The minimum threshold below which the step count intervals will not divide the time.
    auto_min: typing.Optional[str]
    # How many times the current time range should be divided to calculate the value, similar to the Max data points query option.
    # For example, if the current visible time range is 30 minutes, then the auto interval groups the data into 30 one-minute increments.
    auto_count: typing.Optional[int]
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

 * <span class="badge builder"></span> [AdHocVariable](./builder-AdHocVariable.md)
 * <span class="badge builder"></span> [ConstantVariable](./builder-ConstantVariable.md)
 * <span class="badge builder"></span> [CustomVariable](./builder-CustomVariable.md)
 * <span class="badge builder"></span> [DatasourceVariable](./builder-DatasourceVariable.md)
 * <span class="badge builder"></span> [IntervalVariable](./builder-IntervalVariable.md)
 * <span class="badge builder"></span> [QueryVariable](./builder-QueryVariable.md)
 * <span class="badge builder"></span> [TextBoxVariable](./builder-TextBoxVariable.md)
