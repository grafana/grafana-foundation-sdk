---
title: <span class="badge builder"></span> IntervalVariable
---
# <span class="badge builder"></span> IntervalVariable

## Constructor

```python
IntervalVariable(name: str)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.IntervalVariableKind
```

### <span class="badge object-method"></span> auto

```python
def auto(auto: bool) -> typing.Self
```

### <span class="badge object-method"></span> auto_count

```python
def auto_count(auto_count: int) -> typing.Self
```

### <span class="badge object-method"></span> auto_min

```python
def auto_min(auto_min: str) -> typing.Self
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

### <span class="badge object-method"></span> label

```python
def label(label: str) -> typing.Self
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
def spec(spec: dashboardv2beta1.IntervalVariableSpec) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [IntervalVariableKind](./object-IntervalVariableKind.md)
