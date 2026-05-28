---
title: <span class="badge builder"></span> Transformation
---
# <span class="badge builder"></span> Transformation

## Constructor

```python
Transformation()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2.TransformationKind
```

### <span class="badge object-method"></span> disabled

Disabled transformations are skipped

```python
def disabled(disabled: bool) -> typing.Self
```

### <span class="badge object-method"></span> filter_val

Optional frame matcher. When missing it will be applied to all results

```python
def filter_val(filter_val: dashboardv2.MatcherConfig) -> typing.Self
```

### <span class="badge object-method"></span> group

The group is the transformation ID

```python
def group(group: str) -> typing.Self
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
def topic(topic: dashboardv2.DataTopic) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [TransformationKind](./object-TransformationKind.md)
