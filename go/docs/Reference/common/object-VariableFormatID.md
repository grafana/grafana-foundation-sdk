---
title: <span class="badge object-type-enum"></span> VariableFormatID
---
# <span class="badge object-type-enum"></span> VariableFormatID

Optional formats for the template variable replace functions

See also https://grafana.com/docs/grafana/latest/dashboards/variables/variable-syntax/#advanced-variable-format-options

## Definition

```go
type VariableFormatID string
const (
	VariableFormatIDLucene VariableFormatID = "lucene"
	VariableFormatIDRaw VariableFormatID = "raw"
	VariableFormatIDRegex VariableFormatID = "regex"
	VariableFormatIDPipe VariableFormatID = "pipe"
	VariableFormatIDDistributed VariableFormatID = "distributed"
	VariableFormatIDCSV VariableFormatID = "csv"
	VariableFormatIDHTML VariableFormatID = "html"
	VariableFormatIDJSON VariableFormatID = "json"
	VariableFormatIDPercentEncode VariableFormatID = "percentencode"
	VariableFormatIDUriEncode VariableFormatID = "uriencode"
	VariableFormatIDSingleQuote VariableFormatID = "singlequote"
	VariableFormatIDDoubleQuote VariableFormatID = "doublequote"
	VariableFormatIDSQLString VariableFormatID = "sqlstring"
	VariableFormatIDDate VariableFormatID = "date"
	VariableFormatIDGlob VariableFormatID = "glob"
	VariableFormatIDText VariableFormatID = "text"
	VariableFormatIDQueryParam VariableFormatID = "queryparam"
)

```
