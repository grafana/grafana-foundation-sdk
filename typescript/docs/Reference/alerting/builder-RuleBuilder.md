---
title: <span class="badge builder"></span> RuleBuilder
---
# <span class="badge builder"></span> RuleBuilder

## Constructor

```typescript
new RuleBuilder(title: string)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> annotations

```typescript
annotations(annotations: Record<string, string>)
```

### <span class="badge object-method"></span> condition

```typescript
condition(condition: string)
```

### <span class="badge object-method"></span> execErrState

```typescript
execErrState(execErrState: "OK" | "Alerting" | "Error")
```

### <span class="badge object-method"></span> folderUID

```typescript
folderUID(folderUID: string)
```

### <span class="badge object-method"></span> forDuration

The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.

Before this time has elapsed, the rule is only considered to be Pending.

```typescript
forDuration(forVal: string)
```

### <span class="badge object-method"></span> id

```typescript
id(id: number)
```

### <span class="badge object-method"></span> isPaused

```typescript
isPaused(isPaused: boolean)
```

### <span class="badge object-method"></span> labels

```typescript
labels(labels: Record<string, string>)
```

### <span class="badge object-method"></span> noDataState

```typescript
noDataState(noDataState: "Alerting" | "NoData" | "OK")
```

### <span class="badge object-method"></span> orgID

```typescript
orgID(orgID: number)
```

### <span class="badge object-method"></span> provenance

```typescript
provenance(provenance: alerting.Provenance)
```

### <span class="badge object-method"></span> queries

```typescript
queries(data: cog.Builder<alerting.Query>[])
```

### <span class="badge object-method"></span> ruleGroup

```typescript
ruleGroup(ruleGroup: string)
```

### <span class="badge object-method"></span> title

```typescript
title(title: string)
```

### <span class="badge object-method"></span> uid

```typescript
uid(uid: string)
```

### <span class="badge object-method"></span> updated

```typescript
updated(updated: string)
```

### <span class="badge object-method"></span> withQuery

```typescript
withQuery(data: cog.Builder<alerting.Query>)
```

## See also

 * <span class="badge object-type-interface"></span> [Rule](./object-Rule.md)
