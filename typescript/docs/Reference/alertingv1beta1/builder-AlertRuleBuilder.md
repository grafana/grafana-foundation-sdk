---
title: <span class="badge builder"></span> AlertRuleBuilder
---
# <span class="badge builder"></span> AlertRuleBuilder

## Constructor

```typescript
new AlertRuleBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> annotations

```typescript
annotations(annotations: Record<string, alertingv1beta1.TemplateString>)
```

### <span class="badge object-method"></span> execErrState

```typescript
execErrState(execErrState: alertingv1beta1.ExecErrState)
```

### <span class="badge object-method"></span> expressions

```typescript
expressions(expressions: alertingv1beta1.ExpressionMap)
```

### <span class="badge object-method"></span> forVal

```typescript
forVal(forVal: string)
```

### <span class="badge object-method"></span> keepFiringFor

```typescript
keepFiringFor(keepFiringFor: string)
```

### <span class="badge object-method"></span> labels

```typescript
labels(labels: Record<string, alertingv1beta1.TemplateString>)
```

### <span class="badge object-method"></span> missingSeriesEvalsToResolve

```typescript
missingSeriesEvalsToResolve(missingSeriesEvalsToResolve: number)
```

### <span class="badge object-method"></span> noDataState

```typescript
noDataState(noDataState: alertingv1beta1.NoDataState)
```

### <span class="badge object-method"></span> notificationSettings

```typescript
notificationSettings(notificationSettings: cog.Builder<alertingv1beta1.NotificationSettings>)
```

### <span class="badge object-method"></span> panelRef

```typescript
panelRef(panelRef: cog.Builder<alertingv1beta1.PanelRef>)
```

### <span class="badge object-method"></span> paused

```typescript
paused(paused: boolean)
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

 * <span class="badge object-type-interface"></span> [AlertRuleSpec](./object-AlertRuleSpec.md)
