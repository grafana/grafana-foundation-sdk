---
title: <span class="badge object-type-enum"></span> VariableFormatID
---
# <span class="badge object-type-enum"></span> VariableFormatID

Optional formats for the template variable replace functions

See also https://grafana.com/docs/grafana/latest/dashboards/variables/variable-syntax/#advanced-variable-format-options

## Definition

```typescript
export enum VariableFormatID {
	Lucene = "lucene",
	Raw = "raw",
	Regex = "regex",
	Pipe = "pipe",
	Distributed = "distributed",
	CSV = "csv",
	HTML = "html",
	JSON = "json",
	PercentEncode = "percentencode",
	UriEncode = "uriencode",
	SingleQuote = "singlequote",
	DoubleQuote = "doublequote",
	SQLString = "sqlstring",
	Date = "date",
	Glob = "glob",
	Text = "text",
	QueryParam = "queryparam",
}

```
