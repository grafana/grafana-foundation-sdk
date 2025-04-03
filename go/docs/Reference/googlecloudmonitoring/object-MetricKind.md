---
title: <span class="badge object-type-enum"></span> MetricKind
---
# <span class="badge object-type-enum"></span> MetricKind

## Definition

```go
type MetricKind string
const (
	MetricKindMETRICKINDUNSPECIFIED MetricKind = "METRIC_KIND_UNSPECIFIED"
	MetricKindGAUGE MetricKind = "GAUGE"
	MetricKindDELTA MetricKind = "DELTA"
	MetricKindCUMULATIVE MetricKind = "CUMULATIVE"
)

```
