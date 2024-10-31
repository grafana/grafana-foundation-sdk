---
title: <span class="badge object-type-enum"></span> TraceqlSearchScope
---
# <span class="badge object-type-enum"></span> TraceqlSearchScope

static fields are pre-set in the UI, dynamic fields are added by the user

## Definition

```go
type TraceqlSearchScope string
const (
	TraceqlSearchScopeIntrinsic TraceqlSearchScope = "intrinsic"
	TraceqlSearchScopeUnscoped TraceqlSearchScope = "unscoped"
	TraceqlSearchScopeResource TraceqlSearchScope = "resource"
	TraceqlSearchScopeSpan TraceqlSearchScope = "span"
)

```
