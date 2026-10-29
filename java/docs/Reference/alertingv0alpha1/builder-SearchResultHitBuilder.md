---
title: <span class="badge builder"></span> SearchResultHitBuilder
---
# <span class="badge builder"></span> SearchResultHitBuilder

#SearchResultHit is a single match: its identity, an optional relevance

score (present only when the query included free text), and the requested

fields.

## Constructor

```java
new SearchResultHitBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public SearchResultHit build()
```

### <span class="badge object-method"></span> fields

```java
public SearchResultHitBuilder fields(com.grafana.foundation.cog.Builder<RuleSearchHitFields> fields)
```

### <span class="badge object-method"></span> resource

```java
public SearchResultHitBuilder resource(com.grafana.foundation.cog.Builder<SearchResultResource> resource)
```

### <span class="badge object-method"></span> score

```java
public SearchResultHitBuilder score(Double score)
```

## See also

 * <span class="badge object-type-class"></span> [SearchResultHit](./object-SearchResultHit.md)
