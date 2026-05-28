---
title: <span class="badge object-type-enum"></span> FieldColorSeriesByMode
---
# <span class="badge object-type-enum"></span> FieldColorSeriesByMode

Defines how to assign a series color from "by value" color schemes. For example for an aggregated data points like a timeseries, the color can be assigned by the min, max or last value.

## Definition

```go
type FieldColorSeriesByMode string
const (
	FieldColorSeriesByModeMin FieldColorSeriesByMode = "min"
	FieldColorSeriesByModeMax FieldColorSeriesByMode = "max"
	FieldColorSeriesByModeLast FieldColorSeriesByMode = "last"
)

```
