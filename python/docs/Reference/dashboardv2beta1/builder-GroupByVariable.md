---
title: <span class="badge builder"></span> GroupByVariable
---
# <span class="badge builder"></span> GroupByVariable

## Constructor

```python
GroupByVariable(name: str)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.GroupByVariableKind
```

### <span class="badge object-method"></span> current

```python
def current(current: dashboardv2beta1.VariableOption) -> typing.Self
```

### <span class="badge object-method"></span> datasource

```python
def datasource(datasource: cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1GroupByVariableKindDatasource]) -> typing.Self
```

### <span class="badge object-method"></span> default_value

```python
def default_value(default_value: dashboardv2beta1.VariableOption) -> typing.Self
```

### <span class="badge object-method"></span> description

```python
def description(description: str) -> typing.Self
```

### <span class="badge object-method"></span> group

```python
def group(group: str) -> typing.Self
```

### <span class="badge object-method"></span> hide

```python
def hide(hide: dashboardv2beta1.VariableHide) -> typing.Self
```

### <span class="badge object-method"></span> label

```python
def label(label: str) -> typing.Self
```

### <span class="badge object-method"></span> multi

```python
def multi(multi: bool) -> typing.Self
```

### <span class="badge object-method"></span> name

```python
def name(name: str) -> typing.Self
```

### <span class="badge object-method"></span> options

```python
def options(options: list[dashboardv2beta1.VariableOption]) -> typing.Self
```

### <span class="badge object-method"></span> skip_url_sync

```python
def skip_url_sync(skip_url_sync: bool) -> typing.Self
```

### <span class="badge object-method"></span> spec

```python
def spec(spec: dashboardv2beta1.GroupByVariableSpec) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [GroupByVariableKind](./object-GroupByVariableKind.md)
