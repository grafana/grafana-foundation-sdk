---
title: <span class="badge object-type-enum"></span> DataTopic
---
# <span class="badge object-type-enum"></span> DataTopic

A topic is attached to DataFrame metadata in query results.

This specifies where the data should be used.

## Definition

```python
class DataTopic(enum.StrEnum):
    """
    A topic is attached to DataFrame metadata in query results.
    This specifies where the data should be used.
    """

    SERIES = "series"
    ANNOTATIONS = "annotations"
    ALERT_STATES = "alertStates"
```
