---
title: <span class="badge builder"></span> SearchResults
---
# <span class="badge builder"></span> SearchResults

#SearchResults is the response envelope, mirroring

search.grafana.app SearchResults.

## Constructor

```python
SearchResults()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alertingv0alpha1.SearchResults
```

### <span class="badge object-method"></span> facets

```python
def facets(facets: dict[str, list[cogbuilder.Builder[alertingv0alpha1.FacetValue]]]) -> typing.Self
```

### <span class="badge object-method"></span> items

```python
def items(items: list[cogbuilder.Builder[alertingv0alpha1.SearchResultHit]]) -> typing.Self
```

### <span class="badge object-method"></span> metadata

```python
def metadata(metadata: cogbuilder.Builder[alertingv0alpha1.SearchResultsMetadata]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [SearchResults](./object-SearchResults.md)
