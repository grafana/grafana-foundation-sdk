---
title: <span class="badge object-type-enum"></span> QueryType
---
# <span class="badge object-type-enum"></span> QueryType

Defines the supported queryTypes.

## Definition

```python
class QueryType(enum.StrEnum):
    """
    Defines the supported queryTypes.
    """

    TIMESERIESLIST = "timeSeriesList"
    TIMESERIESQUERY = "timeSeriesQuery"
    SLO = "slo"
    ANNOTATION = "annotation"
    PROMQL = "promQL"
```
