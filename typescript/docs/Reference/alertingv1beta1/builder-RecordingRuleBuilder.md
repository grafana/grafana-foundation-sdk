---
title: <span class="badge builder"></span> RecordingRuleBuilder
---
# <span class="badge builder"></span> RecordingRuleBuilder

## Constructor

```typescript
new RecordingRuleBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> expressions

```typescript
expressions(expressions: alertingv1beta1.ExpressionMap)
```

### <span class="badge object-method"></span> labels

```typescript
labels(labels: Record<string, alertingv1beta1.TemplateString>)
```

### <span class="badge object-method"></span> metric

```typescript
metric(metric: alertingv1beta1.MetricName)
```

### <span class="badge object-method"></span> paused

```typescript
paused(paused: boolean)
```

### <span class="badge object-method"></span> targetDatasourceUID

```typescript
targetDatasourceUID(targetDatasourceUID: alertingv1beta1.DatasourceUID)
```

### <span class="badge object-method"></span> title

```typescript
title(title: string)
```

### <span class="badge object-method"></span> trigger

```typescript
trigger(trigger: cog.Builder<alertingv1beta1.IntervalTrigger>)
```

## See also

 * <span class="badge object-type-interface"></span> [RecordingRuleSpec](./object-RecordingRuleSpec.md)
