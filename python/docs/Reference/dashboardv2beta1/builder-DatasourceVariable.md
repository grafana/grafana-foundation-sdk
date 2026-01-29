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
def build() -> dashboardv2beta1.DatasourceVariableKind
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

### <span class="badge object-method"></span> plugin_id

```python
def plugin_id(plugin_id: str) -> typing.Self
```

### <span class="badge object-method"></span> refresh

```python
def refresh(refresh: dashboardv2beta1.VariableRefresh) -> typing.Self
```

### <span class="badge object-method"></span> regex

```python
def regex(regex: str) -> typing.Self
```

### <span class="badge object-method"></span> skip_url_sync

```python
def skip_url_sync(skip_url_sync: bool) -> typing.Self
```

### <span class="badge object-method"></span> spec

```python
def spec(spec: dashboardv2beta1.DatasourceVariableSpec) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [DatasourceVariableKind](./object-DatasourceVariableKind.md)
