---
title: <span class="badge object-type-enum"></span> ThresholdsMode
---
# <span class="badge object-type-enum"></span> ThresholdsMode

Thresholds can either be `absolute` (specific number) or `percentage` (relative to min or max, it will be values between 0 and 1).

## Definition

```python
class ThresholdsMode(enum.StrEnum):
    """
    Thresholds can either be `absolute` (specific number) or `percentage` (relative to min or max, it will be values between 0 and 1).
    """

    ABSOLUTE = "absolute"
    PERCENTAGE = "percentage"
```
