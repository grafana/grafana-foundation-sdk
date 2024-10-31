---
title: <span class="badge object-type-enum"></span> LoadingState
---
# <span class="badge object-type-enum"></span> LoadingState

Loading status

Accepted values are `NotStarted` (the request is not started), `Loading` (waiting for response), `Streaming` (pulling continuous data), `Done` (response received successfully) or `Error` (failed request).

## Definition

```python
class LoadingState(enum.StrEnum):
    """
    Loading status
    Accepted values are `NotStarted` (the request is not started), `Loading` (waiting for response), `Streaming` (pulling continuous data), `Done` (response received successfully) or `Error` (failed request).
    """

    NOT_STARTED = "NotStarted"
    LOADING = "Loading"
    STREAMING = "Streaming"
    DONE = "Done"
    ERROR = "Error"
```
