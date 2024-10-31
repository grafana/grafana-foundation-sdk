---
title: <span class="badge object-type-enum"></span> VariableFormatID
---
# <span class="badge object-type-enum"></span> VariableFormatID

Optional formats for the template variable replace functions

See also https://grafana.com/docs/grafana/latest/dashboards/variables/variable-syntax/#advanced-variable-format-options

## Definition

```python
class VariableFormatID(enum.StrEnum):
    """
    Optional formats for the template variable replace functions
    See also https://grafana.com/docs/grafana/latest/dashboards/variables/variable-syntax/#advanced-variable-format-options
    """

    LUCENE = "lucene"
    RAW = "raw"
    REGEX = "regex"
    PIPE = "pipe"
    DISTRIBUTED = "distributed"
    CSV = "csv"
    HTML = "html"
    JSON = "json"
    PERCENT_ENCODE = "percentencode"
    URI_ENCODE = "uriencode"
    SINGLE_QUOTE = "singlequote"
    DOUBLE_QUOTE = "doublequote"
    SQL_STRING = "sqlstring"
    DATE = "date"
    GLOB = "glob"
    TEXT = "text"
    QUERY_PARAM = "queryparam"
```
