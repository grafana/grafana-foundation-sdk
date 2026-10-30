---
title: <span class="badge builder"></span> AlertRuleInstanceTotalsBuilder
---
# <span class="badge builder"></span> AlertRuleInstanceTotalsBuilder

Count of alert instances per state. error also counts instances whose evaluation

errored but were mapped to another state via execErrState, so it can overlap.

## Constructor

```typescript
new AlertRuleInstanceTotalsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> error

```typescript
error(error: number)
```

### <span class="badge object-method"></span> firing

```typescript
firing(firing: number)
```

### <span class="badge object-method"></span> healthy

```typescript
healthy(healthy: number)
```

### <span class="badge object-method"></span> nodata

```typescript
nodata(nodata: number)
```

### <span class="badge object-method"></span> pending

```typescript
pending(pending: number)
```

### <span class="badge object-method"></span> recovering

```typescript
recovering(recovering: number)
```

## See also

 * <span class="badge object-type-interface"></span> [AlertRuleInstanceTotals](./object-AlertRuleInstanceTotals.md)
