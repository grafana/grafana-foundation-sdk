---
title: <span class="badge builder"></span> Options
---
# <span class="badge builder"></span> Options

## Constructor

```python
Options()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> histogram.Options
```

### <span class="badge object-method"></span> bucket_count

Bucket count (approx)

```python
def bucket_count(bucket_count: int) -> typing.Self
```

### <span class="badge object-method"></span> bucket_offset

Offset buckets by this amount

```python
def bucket_offset(bucket_offset: float) -> typing.Self
```

### <span class="badge object-method"></span> bucket_size

Size of each bucket

```python
def bucket_size(bucket_size: int) -> typing.Self
```

### <span class="badge object-method"></span> combine

Combines multiple series into a single histogram

```python
def combine(combine: bool) -> typing.Self
```

### <span class="badge object-method"></span> legend

```python
def legend(legend: cogbuilder.Builder[common.VizLegendOptions]) -> typing.Self
```

### <span class="badge object-method"></span> tooltip

```python
def tooltip(tooltip: cogbuilder.Builder[common.VizTooltipOptions]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
