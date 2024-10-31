---
title: <span class="badge builder"></span> AnnotationQuery
---
# <span class="badge builder"></span> AnnotationQuery

## Constructor

```python
AnnotationQuery()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboard.AnnotationQuery
```

### <span class="badge object-method"></span> datasource

Datasource where the annotations data is

```python
def datasource(datasource: dashboard.DataSourceRef) -> typing.Self
```

### <span class="badge object-method"></span> enable

When enabled the annotation query is issued with every dashboard refresh

```python
def enable(enable: bool) -> typing.Self
```

### <span class="badge object-method"></span> expr

```python
def expr(expr: str) -> typing.Self
```

### <span class="badge object-method"></span> filter_val

Filters to apply when fetching annotations

```python
def filter_val(filter_val: cogbuilder.Builder[dashboard.AnnotationPanelFilter]) -> typing.Self
```

### <span class="badge object-method"></span> hide

Annotation queries can be toggled on or off at the top of the dashboard.

When hide is true, the toggle is not shown in the dashboard.

```python
def hide(hide: bool) -> typing.Self
```

### <span class="badge object-method"></span> icon_color

Color to use for the annotation event markers

```python
def icon_color(icon_color: str) -> typing.Self
```

### <span class="badge object-method"></span> name

Name of annotation.

```python
def name(name: str) -> typing.Self
```

### <span class="badge object-method"></span> target

TODO.. this should just be a normal query target

```python
def target(target: cogbuilder.Builder[dashboard.AnnotationTarget]) -> typing.Self
```

### <span class="badge object-method"></span> type_val

TODO -- this should not exist here, it is based on the --grafana-- datasource

```python
def type_val(type_val: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [AnnotationQuery](./object-AnnotationQuery.md)
