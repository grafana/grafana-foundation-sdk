---
title: <span class="badge object-type-enum"></span> TempoQueryType
---
# <span class="badge object-type-enum"></span> TempoQueryType

search = Loki search, nativeSearch = Tempo search for backwards compatibility

## Definition

```python
class TempoQueryType(enum.StrEnum):
    """
    search = Loki search, nativeSearch = Tempo search for backwards compatibility
    """

    TRACEQL = "traceql"
    TRACEQL_SEARCH = "traceqlSearch"
    SEARCH = "search"
    SERVICE_MAP = "serviceMap"
    UPLOAD = "upload"
    NATIVE_SEARCH = "nativeSearch"
    TRACE_ID = "traceId"
    CLEAR = "clear"
```
