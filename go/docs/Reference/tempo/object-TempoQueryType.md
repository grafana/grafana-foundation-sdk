---
title: <span class="badge object-type-enum"></span> TempoQueryType
---
# <span class="badge object-type-enum"></span> TempoQueryType

search = Loki search, nativeSearch = Tempo search for backwards compatibility

## Definition

```go
type TempoQueryType string
const (
	TempoQueryTypeTraceql TempoQueryType = "traceql"
	TempoQueryTypeTraceqlSearch TempoQueryType = "traceqlSearch"
	TempoQueryTypeSearch TempoQueryType = "search"
	TempoQueryTypeServiceMap TempoQueryType = "serviceMap"
	TempoQueryTypeUpload TempoQueryType = "upload"
	TempoQueryTypeNativeSearch TempoQueryType = "nativeSearch"
	TempoQueryTypeTraceId TempoQueryType = "traceId"
	TempoQueryTypeClear TempoQueryType = "clear"
)

```
