---
title: <span class="badge builder"></span> RuleSpecBuilder
---
# <span class="badge builder"></span> RuleSpecBuilder

## Constructor

```typescript
new RuleSpecBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> expressions

```typescript
expressions(expressions: alertingv0alpha1.ExpressionMap)
```

### <span class="badge object-method"></span> labels

```typescript
labels(labels: Record<string, alertingv0alpha1.TemplateString>)
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
trigger(trigger: cog.Builder<alertingv0alpha1.IntervalTrigger>)
```

## See also

 * <span class="badge object-type-interface"></span> [RuleSpec](./object-RuleSpec.md)
