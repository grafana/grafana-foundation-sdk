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

### <span class="badge object-method"></span> built_in

Set to 1 for the standard annotation query all dashboards have by default.

```python
def built_in(built_in: float) -> typing.Self
```

### <span class="badge object-method"></span> datasource

Datasource where the annotations data is

```python
def datasource(datasource: common.DataSourceRef) -> typing.Self
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

### <span class="badge object-method"></span> mappings

Mappings define how to convert data frame fields to annotation event fields.

```python
def mappings(mappings: dict[str, cogbuilder.Builder[dashboard.AnnotationEventFieldMapping]]) -> typing.Self
```

### <span class="badge object-method"></span> name

Name of annotation.

```python
def name(name: str) -> typing.Self
```

### <span class="badge object-method"></span> placement

Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.

```python
def placement(placement: str) -> typing.Self
```

### <span class="badge object-method"></span> step

Legacy Prometheus annotation query step interval.

```python
def step(step: str) -> typing.Self
```

### <span class="badge object-method"></span> tag_keys

Comma-separated label keys used as annotation tags.

```python
def tag_keys(tag_keys: str) -> typing.Self
```

### <span class="badge object-method"></span> target

TODO.. this should just be a normal query target

```python
def target(target: cogbuilder.Builder[cogvariants.Dataquery]) -> typing.Self
```

### <span class="badge object-method"></span> text_format

Format for Prometheus annotation text. Label values can be interpolated with templates like {{instance}}.

```python
def text_format(text_format: str) -> typing.Self
```

### <span class="badge object-method"></span> title_format

Format for Prometheus and Loki annotation titles. Label values can be interpolated with templates like {{instance}}.

```python
def title_format(title_format: str) -> typing.Self
```

### <span class="badge object-method"></span> type_val

TODO -- this should not exist here, it is based on the --grafana-- datasource

```python
def type_val(type_val: str) -> typing.Self
```

### <span class="badge object-method"></span> use_value_for_time

Use the Prometheus series value as the annotation timestamp.

```python
def use_value_for_time(use_value_for_time: bool) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [AnnotationQuery](./object-AnnotationQuery.md)
