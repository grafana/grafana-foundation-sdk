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
def build() -> dashboardv2beta1.CustomVariableKind
```

### <span class="badge object-method"></span> all_value

```python
def all_value(all_value: str) -> typing.Self
```

### <span class="badge object-method"></span> allow_custom_value

```python
def allow_custom_value(allow_custom_value: bool) -> typing.Self
```

### <span class="badge object-method"></span> current

```python
def current(current: dashboardv2beta1.VariableOption) -> typing.Self
```

### <span class="badge object-method"></span> description

```python
def description(description: str) -> typing.Self
```

### <span class="badge object-method"></span> hide

```python
def hide(hide: dashboardv2beta1.VariableHide) -> typing.Self
```

### <span class="badge object-method"></span> include_all

```python
def include_all(include_all: bool) -> typing.Self
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

### <span class="badge object-method"></span> query

```python
def query(query: str) -> typing.Self
```

### <span class="badge object-method"></span> skip_url_sync

```python
def skip_url_sync(skip_url_sync: bool) -> typing.Self
```

### <span class="badge object-method"></span> spec

```python
def spec(spec: dashboardv2beta1.CustomVariableSpec) -> typing.Self
```

### <span class="badge object-method"></span> values_format

```python
def values_format(values_format: typing.Literal["csv", "json"]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [CustomVariableKind](./object-CustomVariableKind.md)
