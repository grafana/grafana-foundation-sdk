---
title: <span class="badge builder"></span> DatasourceVariable
---
# <span class="badge builder"></span> DatasourceVariable

## Constructor

```python
DatasourceVariable(name: str)
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

### <span class="badge object-method"></span> all_value

Custom all value

```python
def all_value(all_value: str) -> typing.Self
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

### <span class="badge object-method"></span> include_all

Whether all value option is available or not

```python
def include_all(include_all: bool) -> typing.Self
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

### <span class="badge object-method"></span> regex

Optional field, if you want to extract part of a series name or metric node segment.

Named capture groups can be used to separate the display text and value.

```python
def regex(regex: str) -> typing.Self
```

### <span class="badge object-method"></span> type_val

Query used to fetch values for a variable

```python
def type_val(query: typing.Union[str, dict[str, object]]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [VariableModel](./object-VariableModel.md)
