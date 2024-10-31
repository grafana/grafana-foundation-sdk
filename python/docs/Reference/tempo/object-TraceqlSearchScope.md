---
title: <span class="badge object-type-enum"></span> TraceqlSearchScope
---
# <span class="badge object-type-enum"></span> TraceqlSearchScope

static fields are pre-set in the UI, dynamic fields are added by the user

## Definition

```python
class TraceqlSearchScope(enum.StrEnum):
    """
    static fields are pre-set in the UI, dynamic fields are added by the user
    """

    INTRINSIC = "intrinsic"
    UNSCOPED = "unscoped"
    EVENT = "event"
    INSTRUMENTATION = "instrumentation"
    LINK = "link"
    RESOURCE = "resource"
    SPAN = "span"
```
