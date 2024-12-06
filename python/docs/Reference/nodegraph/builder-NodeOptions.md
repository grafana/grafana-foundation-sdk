---
title: <span class="badge builder"></span> NodeOptions
---
# <span class="badge builder"></span> NodeOptions

## Constructor

```python
NodeOptions()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> nodegraph.NodeOptions
```

### <span class="badge object-method"></span> arcs

Define which fields are shown as part of the node arc (colored circle around the node).

```python
def arcs(arcs: list[cogbuilder.Builder[nodegraph.ArcOption]]) -> typing.Self
```

### <span class="badge object-method"></span> main_stat_unit

Unit for the main stat to override what ever is set in the data frame.

```python
def main_stat_unit(main_stat_unit: str) -> typing.Self
```

### <span class="badge object-method"></span> secondary_stat_unit

Unit for the secondary stat to override what ever is set in the data frame.

```python
def secondary_stat_unit(secondary_stat_unit: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [NodeOptions](./object-NodeOptions.md)
