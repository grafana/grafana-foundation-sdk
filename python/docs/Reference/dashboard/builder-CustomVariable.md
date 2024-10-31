---
title: <span class="badge builder"></span> CustomVariable
---
# <span class="badge builder"></span> CustomVariable

## Constructor

```python
CustomVariable(name: str)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboard.VariableModel
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

### <span class="badge object-method"></span> label

Optional display name

```python
def label(label: str) -> typing.Self
```

### <span class="badge object-method"></span> multi

Whether multiple values can be selected or not from variable value list

```python
def multi(multi: bool) -> typing.Self
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

### <span class="badge object-method"></span> values

Query used to fetch values for a variable

```python
def values(query: typing.Union[str, dict[str, object]]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [VariableModel](./object-VariableModel.md)
