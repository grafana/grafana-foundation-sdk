---
title: <span class="badge builder"></span> ConditionalRenderingGroupBuilder
---
# <span class="badge builder"></span> ConditionalRenderingGroupBuilder

## Constructor

```typescript
new ConditionalRenderingGroupBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> condition

```typescript
condition(condition: "and" | "or")
```

### <span class="badge object-method"></span> item

```typescript
item(item: cog.Builder<dashboardv2beta1.ConditionalRenderingVariableKind> | cog.Builder<dashboardv2beta1.ConditionalRenderingDataKind> | cog.Builder<dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind>)
```

### <span class="badge object-method"></span> items

```typescript
items(items: (cog.Builder<dashboardv2beta1.ConditionalRenderingVariableKind> | cog.Builder<dashboardv2beta1.ConditionalRenderingDataKind> | cog.Builder<dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind>)[])
```

### <span class="badge object-method"></span> visibility

```typescript
visibility(visibility: "show" | "hide")
```

## See also

 * <span class="badge object-type-interface"></span> [ConditionalRenderingGroupKind](./object-ConditionalRenderingGroupKind.md)
