---
title: <span class="badge builder"></span> AnnotationQueryBuilder
---
# <span class="badge builder"></span> AnnotationQueryBuilder

## Constructor

```java
new AnnotationQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public AnnotationQuery build()
```

### <span class="badge object-method"></span> builtIn

```java
public AnnotationQueryBuilder builtIn(Boolean builtIn)
```

### <span class="badge object-method"></span> enable

```java
public AnnotationQueryBuilder enable(Boolean enable)
```

### <span class="badge object-method"></span> filter

```java
public AnnotationQueryBuilder filter(com.grafana.foundation.cog.Builder<AnnotationPanelFilter> filter)
```

### <span class="badge object-method"></span> hide

```java
public AnnotationQueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> iconColor

```java
public AnnotationQueryBuilder iconColor(String iconColor)
```

### <span class="badge object-method"></span> legacyOptions

Catch-all field for datasource-specific properties. Should not be available in as code tooling.

```java
public AnnotationQueryBuilder legacyOptions(Map<String, Object> legacyOptions)
```

### <span class="badge object-method"></span> mappings

Mappings define how to convert data frame fields to annotation event fields.

```java
public AnnotationQueryBuilder mappings(Map<String, com.grafana.foundation.cog.Builder<AnnotationEventFieldMapping>> mappings)
```

### <span class="badge object-method"></span> name

```java
public AnnotationQueryBuilder name(String name)
```

### <span class="badge object-method"></span> placement

Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.

```java
public AnnotationQueryBuilder placement(String placement)
```

### <span class="badge object-method"></span> query

```java
public AnnotationQueryBuilder query(com.grafana.foundation.cog.Builder<DataQueryKind> query)
```

### <span class="badge object-method"></span> spec

```java
public AnnotationQueryBuilder spec(com.grafana.foundation.cog.Builder<AnnotationQuerySpec> spec)
```

## See also

 * <span class="badge object-type-class"></span> [AnnotationQueryKind](./object-AnnotationQueryKind.md)
