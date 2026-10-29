---
title: <span class="badge builder"></span> SearchResultHit
---
# <span class="badge builder"></span> SearchResultHit

#SearchResultHit is a single match: its identity, an optional relevance

score (present only when the query included free text), and the requested

fields.

## Constructor

```python
SearchResultHit()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alertingv0alpha1.SearchResultHit
```

### <span class="badge object-method"></span> fields

```python
def fields(fields: cogbuilder.Builder[alertingv0alpha1.RuleSearchHitFields]) -> typing.Self
```

### <span class="badge object-method"></span> resource

```python
def resource(resource: cogbuilder.Builder[alertingv0alpha1.SearchResultResource]) -> typing.Self
```

### <span class="badge object-method"></span> score

```python
def score(score: float) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [SearchResultHit](./object-SearchResultHit.md)
