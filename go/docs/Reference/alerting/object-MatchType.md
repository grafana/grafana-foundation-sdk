---
title: <span class="badge object-type-enum"></span> MatchType
---
# <span class="badge object-type-enum"></span> MatchType

## Definition

```go
type MatchType string
const (
	MatchTypeEqual MatchType = "="
	MatchTypeNotEqual MatchType = "!="
	MatchTypeEqualRegex MatchType = "=~"
	MatchTypeNotEqualRegex MatchType = "!~"
)

```
