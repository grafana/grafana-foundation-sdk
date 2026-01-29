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
def build() -> dashboardv2beta1.QueryVariableKind
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

### <span class="badge object-method"></span> definition

```python
def definition(definition: str) -> typing.Self
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

### <span class="badge object-method"></span> placeholder

```python
def placeholder(placeholder: str) -> typing.Self
```

### <span class="badge object-method"></span> query

```python
def query(query: cogbuilder.Builder[dashboardv2beta1.DataQueryKind]) -> typing.Self
```

### <span class="badge object-method"></span> refresh

```python
def refresh(refresh: dashboardv2beta1.VariableRefresh) -> typing.Self
```

### <span class="badge object-method"></span> regex

```python
def regex(regex: str) -> typing.Self
```

### <span class="badge object-method"></span> regex_apply_to

```python
def regex_apply_to(regex_apply_to: dashboardv2beta1.VariableRegexApplyTo) -> typing.Self
```

### <span class="badge object-method"></span> skip_url_sync

```python
def skip_url_sync(skip_url_sync: bool) -> typing.Self
```

### <span class="badge object-method"></span> sort

```python
def sort(sort: dashboardv2beta1.VariableSort) -> typing.Self
```

### <span class="badge object-method"></span> spec

```python
def spec(spec: dashboardv2beta1.QueryVariableSpec) -> typing.Self
```

### <span class="badge object-method"></span> static_options

```python
def static_options(static_options: list[dashboardv2beta1.VariableOption]) -> typing.Self
```

### <span class="badge object-method"></span> static_options_order

```python
def static_options_order(static_options_order: typing.Literal["before", "after", "sorted"]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [QueryVariableKind](./object-QueryVariableKind.md)
