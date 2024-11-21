---
title: <span class="badge builder"></span> IntervalVariable
---
# <span class="badge builder"></span> IntervalVariable

## Constructor

```python
IntervalVariable(name: str)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboard.VariableModel
```

### <span class="badge object-method"></span> all_format

Format to use while fetching all values from data source, eg: wildcard, glob, regex, pipe, etc.

```python
def all_format(all_format: str) -> typing.Self
```

### <span class="badge object-method"></span> auto

Dynamically calculates interval by dividing time range by the count specified.

```python
def auto(auto: bool) -> typing.Self
```

### <span class="badge object-method"></span> current

Shows current selected variable text/value on the dashboard

```python
def current(current: dashboard.VariableOption) -> typing.Self
```

### <span class="badge object-method"></span> description

Description of variable. It can be defined but `null`.

```python
def description(description: str) -> typing.Self
```

### <span class="badge object-method"></span> hide

Visibility configuration for the variable

```python
def hide(hide: dashboard.VariableHide) -> typing.Self
```

### <span class="badge object-method"></span> id_val

Unique numeric identifier for the variable.

```python
def id_val(id_val: str) -> typing.Self
```

### <span class="badge object-method"></span> label

Optional display name

```python
def label(label: str) -> typing.Self
```

### <span class="badge object-method"></span> min_interval

The minimum threshold below which the step count intervals will not divide the time.

```python
def min_interval(auto_min: str) -> typing.Self
```

### <span class="badge object-method"></span> name

Name of variable

```python
def name(name: str) -> typing.Self
```

### <span class="badge object-method"></span> options

Options that can be selected for a variable.

```python
def options(options: list[dashboard.VariableOption]) -> typing.Self
```

### <span class="badge object-method"></span> step_count

How many times the current time range should be divided to calculate the value, similar to the Max data points query option.

For example, if the current visible time range is 30 minutes, then the auto interval groups the data into 30 one-minute increments.

```python
def step_count(auto_count: int) -> typing.Self
```

### <span class="badge object-method"></span> values

Query used to fetch values for a variable

```python
def values(query: typing.Union[str, dict[str, object]]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [VariableModel](./object-VariableModel.md)
