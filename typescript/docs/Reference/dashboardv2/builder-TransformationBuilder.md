---
title: <span class="badge builder"></span> TransformationBuilder
---
# <span class="badge builder"></span> TransformationBuilder

## Constructor

```typescript
new TransformationBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> disabled

Disabled transformations are skipped

```typescript
disabled(disabled: boolean)
```

### <span class="badge object-method"></span> filter

Optional frame matcher. When missing it will be applied to all results

```typescript
filter(filter: dashboardv2.MatcherConfig)
```

### <span class="badge object-method"></span> group

The group is the transformation ID

```typescript
group(group: string)
```

### <span class="badge object-method"></span> options

Options to be passed to the transformer

Valid options depend on the transformer id

```typescript
options(options: any)
```

### <span class="badge object-method"></span> topic

Where to pull DataFrames from as input to transformation

```typescript
topic(topic: dashboardv2.DataTopic)
```

## See also

 * <span class="badge object-type-interface"></span> [TransformationKind](./object-TransformationKind.md)
