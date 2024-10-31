---
title: <span class="badge object-type-enum"></span> TypeThresholdType
---
# <span class="badge object-type-enum"></span> TypeThresholdType

## Definition

```go
type TypeThresholdType string
const (
	TypeThresholdTypeNone TypeThresholdType = ""
	TypeThresholdTypeTimeseriesWide TypeThresholdType = "timeseries-wide"
	TypeThresholdTypeTimeseriesLong TypeThresholdType = "timeseries-long"
	TypeThresholdTypeTimeseriesMany TypeThresholdType = "timeseries-many"
	TypeThresholdTypeTimeseriesMulti TypeThresholdType = "timeseries-multi"
	TypeThresholdTypeDirectoryListing TypeThresholdType = "directory-listing"
	TypeThresholdTypeTable TypeThresholdType = "table"
	TypeThresholdTypeNumericWide TypeThresholdType = "numeric-wide"
	TypeThresholdTypeNumericMulti TypeThresholdType = "numeric-multi"
	TypeThresholdTypeNumericLong TypeThresholdType = "numeric-long"
	TypeThresholdTypeLogLines TypeThresholdType = "log-lines"
)

```
