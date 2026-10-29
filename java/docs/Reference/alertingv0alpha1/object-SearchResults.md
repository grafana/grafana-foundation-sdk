---
title: <span class="badge object-type-class"></span> SearchResults
---
# <span class="badge object-type-class"></span> SearchResults

#SearchResults is the response envelope, mirroring

search.grafana.app SearchResults.

## Definition

```java
public class SearchResults {
  public SearchResultsMetadata metadata;
  public List<SearchResultHit> items;
  public Map<String, List<FacetValue>> facets;
}
```
## See also

 * <span class="badge builder"></span> [SearchResultsBuilder](./builder-SearchResultsBuilder.md)
