---
title: <span class="badge builder"></span> AnnotationQuerySpecBuilder
---
# <span class="badge builder"></span> AnnotationQuerySpecBuilder

## Constructor

```java
new AnnotationQuerySpecBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public AnnotationQuerySpec build()
```

### <span class="badge object-method"></span> builtIn

```java
public AnnotationQuerySpecBuilder builtIn(Boolean builtIn)
```

### <span class="badge object-method"></span> enable

```java
public AnnotationQuerySpecBuilder enable(Boolean enable)
```

### <span class="badge object-method"></span> filter

```java
public AnnotationQuerySpecBuilder filter(com.grafana.foundation.cog.Builder<AnnotationPanelFilter> filter)
```

### <span class="badge object-method"></span> hide

```java
public AnnotationQuerySpecBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> iconColor

```java
public AnnotationQuerySpecBuilder iconColor(String iconColor)
```

### <span class="badge object-method"></span> legacyOptions

Catch-all field for datasource-specific properties. Should not be available in as code tooling.

```java
public AnnotationQuerySpecBuilder legacyOptions(Map<String, Object> legacyOptions)
```

### <span class="badge object-method"></span> mappings

Mappings define how to convert data frame fields to annotation event fields.

```java
public AnnotationQuerySpecBuilder mappings(Map<String, com.grafana.foundation.cog.Builder<AnnotationEventFieldMapping>> mappings)
```

### <span class="badge object-method"></span> name

```java
public AnnotationQuerySpecBuilder name(String name)
```

### <span class="badge object-method"></span> placement

Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.

```java
public AnnotationQuerySpecBuilder placement(String placement)
```

### <span class="badge object-method"></span> query

```java
public AnnotationQuerySpecBuilder query(com.grafana.foundation.cog.Builder<DataQueryKind> query)
```

## See also

 * <span class="badge object-type-class"></span> [AnnotationQuerySpec](./object-AnnotationQuerySpec.md)
