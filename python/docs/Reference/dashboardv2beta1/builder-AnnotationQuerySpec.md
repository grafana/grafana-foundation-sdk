---
title: <span class="badge builder"></span> AnnotationQuerySpec
---
# <span class="badge builder"></span> AnnotationQuerySpec

## Constructor

```python
AnnotationQuerySpec()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.AnnotationQuerySpec
```

### <span class="badge object-method"></span> built_in

```python
def built_in(built_in: bool) -> typing.Self
```

### <span class="badge object-method"></span> enable

```python
def enable(enable: bool) -> typing.Self
```

### <span class="badge object-method"></span> filter_val

```python
def filter_val(filter_val: cogbuilder.Builder[dashboardv2beta1.AnnotationPanelFilter]) -> typing.Self
```

### <span class="badge object-method"></span> hide

```python
def hide(hide: bool) -> typing.Self
```

### <span class="badge object-method"></span> icon_color

```python
def icon_color(icon_color: str) -> typing.Self
```

### <span class="badge object-method"></span> legacy_options

Catch-all field for datasource-specific properties. Should not be available in as code tooling.

```python
def legacy_options(legacy_options: dict[str, object]) -> typing.Self
```

### <span class="badge object-method"></span> mappings

Mappings define how to convert data frame fields to annotation event fields.

```python
def mappings(mappings: dict[str, cogbuilder.Builder[dashboardv2beta1.AnnotationEventFieldMapping]]) -> typing.Self
```

### <span class="badge object-method"></span> name

```python
def name(name: str) -> typing.Self
```

### <span class="badge object-method"></span> placement

Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.

```python
def placement(placement: str) -> typing.Self
```

### <span class="badge object-method"></span> query

```python
def query(query: cogbuilder.Builder[dashboardv2beta1.DataQueryKind]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [AnnotationQuerySpec](./object-AnnotationQuerySpec.md)
