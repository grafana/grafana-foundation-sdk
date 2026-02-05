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

```typescript
builtIn(builtIn: boolean)
```

### <span class="badge object-method"></span> enable

```typescript
enable(enable: boolean)
```

### <span class="badge object-method"></span> filter

```typescript
filter(filter: cog.Builder<dashboardv2beta1.AnnotationPanelFilter>)
```

### <span class="badge object-method"></span> hide

```typescript
hide(hide: boolean)
```

### <span class="badge object-method"></span> iconColor

```typescript
iconColor(iconColor: string)
```

### <span class="badge object-method"></span> legacyOptions

Catch-all field for datasource-specific properties. Should not be available in as code tooling.

```typescript
legacyOptions(legacyOptions: Record<string, any>)
```

### <span class="badge object-method"></span> name

```typescript
name(name: string)
```

### <span class="badge object-method"></span> placement

Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.

```typescript
placement(placement: "inControlsMenu")
```

### <span class="badge object-method"></span> query

```typescript
query(query: cog.Builder<dashboardv2beta1.DataQueryKind>)
```

### <span class="badge object-method"></span> spec

```typescript
spec(spec: cog.Builder<dashboardv2beta1.AnnotationQuerySpec>)
```

## See also

 * <span class="badge object-type-interface"></span> [AnnotationQueryKind](./object-AnnotationQueryKind.md)
