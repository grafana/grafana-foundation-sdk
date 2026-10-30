---
title: <span class="badge builder"></span> RuleSequenceBuilder
---
# <span class="badge builder"></span> RuleSequenceBuilder

## Constructor

```typescript
new RuleSequenceBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> alertRules

```typescript
alertRules(alertRules: cog.Builder<alertingv1beta1.RuleRef>[])
```

### <span class="badge object-method"></span> recordingRules

```typescript
recordingRules(recordingRules: cog.Builder<alertingv1beta1.RuleRef>[])
```

### <span class="badge object-method"></span> trigger

```typescript
trigger(trigger: cog.Builder<alertingv1beta1.IntervalTrigger>)
```

## See also

 * <span class="badge object-type-interface"></span> [RuleSequenceSpec](./object-RuleSequenceSpec.md)
