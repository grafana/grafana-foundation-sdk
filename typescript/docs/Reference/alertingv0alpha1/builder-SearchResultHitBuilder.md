---
title: <span class="badge builder"></span> SearchResultHitBuilder
---
# <span class="badge builder"></span> SearchResultHitBuilder

#SearchResultHit is a single match: its identity, an optional relevance

score (present only when the query included free text), and the requested

fields.

## Constructor

```typescript
new SearchResultHitBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> fields

```typescript
fields(fields: cog.Builder<alertingv0alpha1.RuleSearchHitFields>)
```

### <span class="badge object-method"></span> resource

```typescript
resource(resource: cog.Builder<alertingv0alpha1.SearchResultResource>)
```

### <span class="badge object-method"></span> score

```typescript
score(score: number)
```

## See also

 * <span class="badge object-type-interface"></span> [SearchResultHit](./object-SearchResultHit.md)
