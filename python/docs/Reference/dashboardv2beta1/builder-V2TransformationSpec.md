---
title: <span class="badge builder"></span> V2TransformationSpec
---
# <span class="badge builder"></span> V2TransformationSpec

Dashboard v2 shape: no `id`, it is carried by the parent's `group`.

## Constructor

```python
V2TransformationSpec()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.V2TransformationSpec
```

### <span class="badge object-method"></span> disabled

Disabled transformations are skipped

```python
def disabled(disabled: bool) -> typing.Self
```

### <span class="badge object-method"></span> filter_val

Optional frame matcher. When missing it will be applied to all results

```python
def filter_val(filter_val: dashboardv2beta1.MatcherConfig) -> typing.Self
```

### <span class="badge object-method"></span> options

Options to be passed to the transformer

Valid options depend on the transformer id

```python
def options(options: object) -> typing.Self
```

### <span class="badge object-method"></span> topic

Where to pull DataFrames from as input to transformation

```python
def topic(topic: dashboardv2beta1.DataTopic) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [V2TransformationSpec](./object-V2TransformationSpec.md)
