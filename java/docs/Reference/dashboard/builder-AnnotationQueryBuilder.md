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

Set to 1 for the standard annotation query all dashboards have by default.

```java
public AnnotationQueryBuilder builtIn(Double builtIn)
```

### <span class="badge object-method"></span> datasource

Datasource where the annotations data is

```java
public AnnotationQueryBuilder datasource(DataSourceRef datasource)
```

### <span class="badge object-method"></span> enable

When enabled the annotation query is issued with every dashboard refresh

```java
public AnnotationQueryBuilder enable(Boolean enable)
```

### <span class="badge object-method"></span> expr

```java
public AnnotationQueryBuilder expr(String expr)
```

### <span class="badge object-method"></span> filter

Filters to apply when fetching annotations

```java
public AnnotationQueryBuilder filter(AnnotationPanelFilter filter)
```

### <span class="badge object-method"></span> hide

Annotation queries can be toggled on or off at the top of the dashboard.

When hide is true, the toggle is not shown in the dashboard.

```java
public AnnotationQueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> iconColor

Color to use for the annotation event markers

```java
public AnnotationQueryBuilder iconColor(String iconColor)
```

### <span class="badge object-method"></span> name

Name of annotation.

```java
public AnnotationQueryBuilder name(String name)
```

### <span class="badge object-method"></span> target

TODO.. this should just be a normal query target

```java
public AnnotationQueryBuilder target(AnnotationTarget target)
```

### <span class="badge object-method"></span> type

TODO -- this should not exist here, it is based on the --grafana-- datasource

```java
public AnnotationQueryBuilder type(String type)
```

## See also

 * <span class="badge object-type-class"></span> [AnnotationQuery](./object-AnnotationQuery.md)
