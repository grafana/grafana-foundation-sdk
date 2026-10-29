---
title: <span class="badge builder"></span> RuleSearchHitFieldsBuilder
---
# <span class="badge builder"></span> RuleSearchHitFieldsBuilder

#RuleSearchHitFields is the per-kind field payload returned on each hit.

It carries the union of alert- and recording-rule search fields; only the

fields relevant to a hit's kind are populated. This maps to the kind's

declared searchFields.

## Constructor

```typescript
new RuleSearchHitFieldsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> annotations

Alert-rule fields.

```typescript
annotations(annotations: Record<string, string>)
```

### <span class="badge object-method"></span> dashboardUID

```typescript
dashboardUID(dashboardUID: string)
```

### <span class="badge object-method"></span> datasourceUIDs

```typescript
datasourceUIDs(datasourceUIDs: string[])
```

### <span class="badge object-method"></span> folder

```typescript
folder(folder: string)
```

### <span class="badge object-method"></span> forVal

```typescript
forVal(forVal: string)
```

### <span class="badge object-method"></span> interval

```typescript
interval(interval: string)
```

### <span class="badge object-method"></span> keepFiringFor

```typescript
keepFiringFor(keepFiringFor: string)
```

### <span class="badge object-method"></span> labels

```typescript
labels(labels: Record<string, string>)
```

### <span class="badge object-method"></span> metric

Recording-rule fields.

```typescript
metric(metric: string)
```

### <span class="badge object-method"></span> notificationType

```typescript
notificationType(notificationType: string)
```

### <span class="badge object-method"></span> panelID

```typescript
panelID(panelID: number)
```

### <span class="badge object-method"></span> paused

```typescript
paused(paused: boolean)
```

### <span class="badge object-method"></span> receiver

```typescript
receiver(receiver: string)
```

### <span class="badge object-method"></span> routingTree

```typescript
routingTree(routingTree: string)
```

### <span class="badge object-method"></span> targetDatasourceUID

```typescript
targetDatasourceUID(targetDatasourceUID: string)
```

### <span class="badge object-method"></span> title

```typescript
title(title: string)
```

### <span class="badge object-method"></span> type

```typescript
type(type: string)
```

## See also

 * <span class="badge object-type-interface"></span> [RuleSearchHitFields](./object-RuleSearchHitFields.md)
