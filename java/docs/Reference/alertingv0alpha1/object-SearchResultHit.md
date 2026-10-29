---
title: <span class="badge object-type-class"></span> SearchResultHit
---
# <span class="badge object-type-class"></span> SearchResultHit

#SearchResultHit is a single match: its identity, an optional relevance

score (present only when the query included free text), and the requested

fields.

## Definition

```java
public class SearchResultHit {
  public SearchResultResource resource;
  public Double score;
  public RuleSearchHitFields fields;
}
```
## See also

 * <span class="badge builder"></span> [SearchResultHitBuilder](./builder-SearchResultHitBuilder.md)
