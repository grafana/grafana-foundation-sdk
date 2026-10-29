---
title: <span class="badge builder"></span> SearchResultsBuilder
---
# <span class="badge builder"></span> SearchResultsBuilder

#SearchResults is the response envelope, mirroring

search.grafana.app SearchResults.

## Constructor

```java
new SearchResultsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public SearchResults build()
```

### <span class="badge object-method"></span> facets

```java
public SearchResultsBuilder facets(Map<String, List<com.grafana.foundation.cog.Builder<FacetValue>>> facets)
```

### <span class="badge object-method"></span> items

```java
public SearchResultsBuilder items(List<com.grafana.foundation.cog.Builder<SearchResultHit>> items)
```

### <span class="badge object-method"></span> metadata

```java
public SearchResultsBuilder metadata(com.grafana.foundation.cog.Builder<SearchResultsMetadata> metadata)
```

## See also

 * <span class="badge object-type-class"></span> [SearchResults](./object-SearchResults.md)
