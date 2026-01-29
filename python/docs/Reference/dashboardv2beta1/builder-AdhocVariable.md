---
title: <span class="badge builder"></span> AdhocVariable
---
# <span class="badge builder"></span> AdhocVariable

## Constructor

```python
AdhocVariable(name: str)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.AdhocVariableKind
```

### <span class="badge object-method"></span> allow_custom_value

```python
def allow_custom_value(allow_custom_value: bool) -> typing.Self
```

### <span class="badge object-method"></span> base_filters

```python
def base_filters(base_filters: list[cogbuilder.Builder[dashboardv2beta1.AdHocFilterWithLabels]]) -> typing.Self
```

### <span class="badge object-method"></span> datasource

```python
def datasource(datasource: cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1AdhocVariableKindDatasource]) -> typing.Self
```

### <span class="badge object-method"></span> default_keys

```python
def default_keys(default_keys: list[cogbuilder.Builder[dashboardv2beta1.MetricFindValue]]) -> typing.Self
```

### <span class="badge object-method"></span> description

```python
def description(description: str) -> typing.Self
```

### <span class="badge object-method"></span> filters

```python
def filters(filters: list[cogbuilder.Builder[dashboardv2beta1.AdHocFilterWithLabels]]) -> typing.Self
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

### <span class="badge object-method"></span> name

```python
def name(name: str) -> typing.Self
```

### <span class="badge object-method"></span> skip_url_sync

```python
def skip_url_sync(skip_url_sync: bool) -> typing.Self
```

### <span class="badge object-method"></span> spec

```python
def spec(spec: dashboardv2beta1.AdhocVariableSpec) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [AdhocVariableKind](./object-AdhocVariableKind.md)
