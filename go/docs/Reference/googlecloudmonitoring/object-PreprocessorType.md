---
title: <span class="badge object-type-enum"></span> PreprocessorType
---
# <span class="badge object-type-enum"></span> PreprocessorType

Types of pre-processor available. Defined by the metric.

## Definition

```go
type PreprocessorType string
const (
	PreprocessorTypeNone PreprocessorType = "none"
	PreprocessorTypeRate PreprocessorType = "rate"
	PreprocessorTypeDelta PreprocessorType = "delta"
)

```
