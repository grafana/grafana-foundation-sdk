---
title: <span class="badge builder"></span> ConstantVariable
---
# <span class="badge builder"></span> ConstantVariable

## Constructor

```python
ConstantVariable(name: str)
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

### <span class="badge object-method"></span> description

Description of variable. It can be defined but `null`.

```python
def description(description: str) -> typing.Self
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

### <span class="badge object-method"></span> name

Name of variable

```python
def name(name: str) -> typing.Self
```

### <span class="badge object-method"></span> value

Query used to fetch values for a variable

```python
def value(query: typing.Union[str, dict[str, object]]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [VariableModel](./object-VariableModel.md)
