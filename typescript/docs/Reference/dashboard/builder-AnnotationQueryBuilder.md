---
title: <span class="badge builder"></span> AnnotationQueryBuilder
---
# <span class="badge builder"></span> AnnotationQueryBuilder

## Constructor

```typescript
new AnnotationQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> builtIn

Set to 1 for the standard annotation query all dashboards have by default.

```typescript
builtIn(builtIn: number)
```

### <span class="badge object-method"></span> datasource

Datasource where the annotations data is

```typescript
datasource(datasource: dashboard.DataSourceRef)
```

### <span class="badge object-method"></span> enable

When enabled the annotation query is issued with every dashboard refresh

```typescript
enable(enable: boolean)
```

### <span class="badge object-method"></span> expr

```typescript
expr(expr: string)
```

### <span class="badge object-method"></span> filter

Filters to apply when fetching annotations

```typescript
filter(filter: cog.Builder<dashboard.AnnotationPanelFilter>)
```

### <span class="badge object-method"></span> hide

Annotation queries can be toggled on or off at the top of the dashboard.

When hide is true, the toggle is not shown in the dashboard.

```typescript
hide(hide: boolean)
```

### <span class="badge object-method"></span> iconColor

Color to use for the annotation event markers

```typescript
iconColor(iconColor: string)
```

### <span class="badge object-method"></span> name

Name of annotation.

```typescript
name(name: string)
```

### <span class="badge object-method"></span> target

TODO.. this should just be a normal query target

```typescript
target(target: cog.Builder<dashboard.AnnotationTarget>)
```

### <span class="badge object-method"></span> type

TODO -- this should not exist here, it is based on the --grafana-- datasource

```typescript
type(type: string)
```

## See also

 * <span class="badge object-type-interface"></span> [AnnotationQuery](./object-AnnotationQuery.md)
