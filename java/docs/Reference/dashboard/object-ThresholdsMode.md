---
title: <span class="badge object-type-enum"></span> ThresholdsMode
---
# <span class="badge object-type-enum"></span> ThresholdsMode

Thresholds can either be `absolute` (specific number) or `percentage` (relative to min or max, it will be values between 0 and 1).

## Definition

```java
package com.grafana.foundation.dashboard.ThresholdsMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Thresholds can either be `absolute` (specific number) or `percentage` (relative to min or max, it will be values between 0 and 1).
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ThresholdsMode {
    ABSOLUTE("absolute"),
    PERCENTAGE("percentage"),
    _EMPTY("");

    private final String value;

    private ThresholdsMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
