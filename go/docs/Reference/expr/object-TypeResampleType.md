---
title: <span class="badge object-type-enum"></span> TypeResampleType
---
# <span class="badge object-type-enum"></span> TypeResampleType

## Definition

```go
type TypeResampleType string
const (
	TypeResampleTypeNone TypeResampleType = ""
	TypeResampleTypeTimeseriesWide TypeResampleType = "timeseries-wide"
	TypeResampleTypeTimeseriesLong TypeResampleType = "timeseries-long"
	TypeResampleTypeTimeseriesMany TypeResampleType = "timeseries-many"
	TypeResampleTypeTimeseriesMulti TypeResampleType = "timeseries-multi"
	TypeResampleTypeDirectoryListing TypeResampleType = "directory-listing"
	TypeResampleTypeTable TypeResampleType = "table"
	TypeResampleTypeNumericWide TypeResampleType = "numeric-wide"
	TypeResampleTypeNumericMulti TypeResampleType = "numeric-multi"
	TypeResampleTypeNumericLong TypeResampleType = "numeric-long"
	TypeResampleTypeLogLines TypeResampleType = "log-lines"
)

```
