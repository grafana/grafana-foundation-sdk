---
title: <span class="badge builder"></span> NotebookSpec
---
# <span class="badge builder"></span> NotebookSpec

A notebook spec is a dashboard spec with narrative content. It has a title, optional description, tags, time

settings, and a map of elements (panels and cells) referenced by the layout.

## Constructor

```python
NotebookSpec()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.NotebookSpec
```

### <span class="badge object-method"></span> description

```python
def description(description: str) -> typing.Self
```

### <span class="badge object-method"></span> elements

```python
def elements(elements: dict[str, cogbuilder.Builder[dashboardv2beta1.NotebookElement]]) -> typing.Self
```

### <span class="badge object-method"></span> layout

```python
def layout(layout: cogbuilder.Builder[dashboardv2beta1.NotebookLayoutKind]) -> typing.Self
```

### <span class="badge object-method"></span> tags

```python
def tags(tags: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> time_settings

```python
def time_settings(time_settings: cogbuilder.Builder[dashboardv2beta1.TimeSettingsSpec]) -> typing.Self
```

### <span class="badge object-method"></span> title

```python
def title(title: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [NotebookSpec](./object-NotebookSpec.md)
