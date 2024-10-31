---
title: <span class="badge builder"></span> AdHocVariable
---
# <span class="badge builder"></span> AdHocVariable

## Constructor

```python
AdHocVariable(name: str)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboard.VariableModel
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

## See also

 * <span class="badge object-type-class"></span> [VariableModel](./object-VariableModel.md)
