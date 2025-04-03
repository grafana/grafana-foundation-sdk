---
title: <span class="badge builder"></span> QueryHistoryPreferenceBuilder
---
# <span class="badge builder"></span> QueryHistoryPreferenceBuilder

## Constructor

```go
func NewQueryHistoryPreferenceBuilder() *QueryHistoryPreferenceBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *QueryHistoryPreferenceBuilder) Build() (QueryHistoryPreference, error)
```

### <span class="badge object-method"></span> HomeTab

one of: '' | 'query' | 'starred';

```go
func (builder *QueryHistoryPreferenceBuilder) HomeTab(homeTab string) *QueryHistoryPreferenceBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [QueryHistoryPreference](./object-QueryHistoryPreference.md)
