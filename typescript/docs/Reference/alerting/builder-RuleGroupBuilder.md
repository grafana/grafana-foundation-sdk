---
title: <span class="badge builder"></span> <span class="badge deprecated"></span> RuleGroupBuilder
---
# <span class="badge builder"></span> <span class="badge deprecated"></span> RuleGroupBuilder

<span class="badge deprecated"></span>Prefer using the alertingv1beta1 package instead

## Constructor

```typescript
new RuleGroupBuilder(title: string)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> folderUid

```typescript
folderUid(folderUid: string)
```

### <span class="badge object-method"></span> interval

The interval, in seconds, at which all rules in the group are evaluated.

If a group contains many rules, the rules are evaluated sequentially.

```typescript
interval(interval: alerting.Duration)
```

### <span class="badge object-method"></span> rules

```typescript
rules(rules: cog.Builder<alerting.Rule>[])
```

### <span class="badge object-method"></span> title

```typescript
title(title: string)
```

### <span class="badge object-method"></span> withRule

```typescript
withRule(rule: cog.Builder<alerting.Rule>)
```

## See also

 * <span class="badge object-type-interface"></span> [RuleGroup](./object-RuleGroup.md)
