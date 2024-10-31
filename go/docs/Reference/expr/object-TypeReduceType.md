---
title: <span class="badge object-type-enum"></span> TypeReduceType
---
# <span class="badge object-type-enum"></span> TypeReduceType

## Definition

```go
type TypeReduceType string
const (
	TypeReduceTypeNone TypeReduceType = ""
	TypeReduceTypeTimeseriesWide TypeReduceType = "timeseries-wide"
	TypeReduceTypeTimeseriesLong TypeReduceType = "timeseries-long"
	TypeReduceTypeTimeseriesMany TypeReduceType = "timeseries-many"
	TypeReduceTypeTimeseriesMulti TypeReduceType = "timeseries-multi"
	TypeReduceTypeDirectoryListing TypeReduceType = "directory-listing"
	TypeReduceTypeTable TypeReduceType = "table"
	TypeReduceTypeNumericWide TypeReduceType = "numeric-wide"
	TypeReduceTypeNumericMulti TypeReduceType = "numeric-multi"
	TypeReduceTypeNumericLong TypeReduceType = "numeric-long"
	TypeReduceTypeLogLines TypeReduceType = "log-lines"
)

```
