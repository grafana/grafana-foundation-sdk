---
title: <span class="badge object-type-enum"></span> QueryType
---
# <span class="badge object-type-enum"></span> QueryType

Defines the supported queryTypes.

## Definition

```go
type QueryType string
const (
	QueryTypeTIMESERIESLIST QueryType = "timeSeriesList"
	QueryTypeTIMESERIESQUERY QueryType = "timeSeriesQuery"
	QueryTypeSLO QueryType = "slo"
	QueryTypeANNOTATION QueryType = "annotation"
)

```
