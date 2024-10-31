---
title: <span class="badge object-type-enum"></span> SearchStreamingState
---
# <span class="badge object-type-enum"></span> SearchStreamingState

The state of the TraceQL streaming search query

## Definition

```python
class SearchStreamingState(enum.StrEnum):
    """
    The state of the TraceQL streaming search query
    """

    PENDING = "pending"
    STREAMING = "streaming"
    DONE = "done"
    ERROR = "error"
```
