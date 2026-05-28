---
title: <span class="badge object-type-enum"></span> FieldColorSeriesByMode
---
# <span class="badge object-type-enum"></span> FieldColorSeriesByMode

Defines how to assign a series color from "by value" color schemes. For example for an aggregated data points like a timeseries, the color can be assigned by the min, max or last value.

## Definition

```python
class FieldColorSeriesByMode(enum.StrEnum):
    """
    Defines how to assign a series color from "by value" color schemes. For example for an aggregated data points like a timeseries, the color can be assigned by the min, max or last value.
    """

    MIN = "min"
    MAX = "max"
    LAST = "last"
```
