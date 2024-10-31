---
title: <span class="badge builder"></span> AzureTracesFilter
---
# <span class="badge builder"></span> AzureTracesFilter

## Constructor

```python
AzureTracesFilter()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> azuremonitor.AzureTracesFilter
```

### <span class="badge object-method"></span> filters

Values to filter by.

```python
def filters(filters: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> operation

Comparison operator to use. Either equals or not equals.

```python
def operation(operation: str) -> typing.Self
```

### <span class="badge object-method"></span> property_val

Property name, auto-populated based on available traces.

```python
def property_val(property_val: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [AzureTracesFilter](./object-AzureTracesFilter.md)
