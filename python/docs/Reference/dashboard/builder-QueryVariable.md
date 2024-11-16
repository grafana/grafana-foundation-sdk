---
title: <span class="badge builder"></span> QueryVariable
---
# <span class="badge builder"></span> QueryVariable

## Constructor

```python
QueryVariable(name: str)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboard.VariableModel
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

### <span class="badge object-method"></span> datasource

Data source used to fetch values for a variable. It can be defined but `null`.

```python
def datasource(datasource: dashboard.DataSourceRef) -> typing.Self
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

### <span class="badge object-method"></span> options

Options that can be selected for a variable.

```python
def options(options: list[dashboard.VariableOption]) -> typing.Self
```

### <span class="badge object-method"></span> query

Query used to fetch values for a variable

```python
def query(query: typing.Union[str, dict[str, object]]) -> typing.Self
```

### <span class="badge object-method"></span> refresh

Options to config when to refresh a variable

```python
def refresh(refresh: dashboard.VariableRefresh) -> typing.Self
```

### <span class="badge object-method"></span> regex

Optional field, if you want to extract part of a series name or metric node segment.

Named capture groups can be used to separate the display text and value.

```python
def regex(regex: str) -> typing.Self
```

### <span class="badge object-method"></span> sort

Options sort order

```python
def sort(sort: dashboard.VariableSort) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [VariableModel](./object-VariableModel.md)