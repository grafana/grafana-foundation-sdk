---
title: <span class="badge builder"></span> AzureMetricDimension
---
# <span class="badge builder"></span> AzureMetricDimension

## Constructor

```python
AzureMetricDimension()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> azuremonitor.AzureMetricDimension
```

### <span class="badge object-method"></span> dimension

Name of Dimension to be filtered on.

```python
def dimension(dimension: str) -> typing.Self
```

### <span class="badge object-method"></span> filter_val

@deprecated filter is deprecated in favour of filters to support multiselect.

```python
def filter_val(filter_val: str) -> typing.Self
```

### <span class="badge object-method"></span> filters

Values to match with the filter.

```python
def filters(filters: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> operator

String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.

```python
def operator(operator: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [AzureMetricDimension](./object-AzureMetricDimension.md)
