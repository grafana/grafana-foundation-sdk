---
title: <span class="badge object-type-enum"></span> ResultAssertionsType
---
# <span class="badge object-type-enum"></span> ResultAssertionsType

## Definition

```go
type ResultAssertionsType string
const (
	ResultAssertionsTypeNone ResultAssertionsType = ""
	ResultAssertionsTypeTimeseriesWide ResultAssertionsType = "timeseries-wide"
	ResultAssertionsTypeTimeseriesLong ResultAssertionsType = "timeseries-long"
	ResultAssertionsTypeTimeseriesMany ResultAssertionsType = "timeseries-many"
	ResultAssertionsTypeTimeseriesMulti ResultAssertionsType = "timeseries-multi"
	ResultAssertionsTypeDirectoryListing ResultAssertionsType = "directory-listing"
	ResultAssertionsTypeTable ResultAssertionsType = "table"
	ResultAssertionsTypeNumericWide ResultAssertionsType = "numeric-wide"
	ResultAssertionsTypeNumericMulti ResultAssertionsType = "numeric-multi"
	ResultAssertionsTypeNumericLong ResultAssertionsType = "numeric-long"
	ResultAssertionsTypeLogLines ResultAssertionsType = "log-lines"
)

```
