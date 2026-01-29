---
title: <span class="badge object-type-enum"></span> ThresholdsMode
---
# <span class="badge object-type-enum"></span> ThresholdsMode

## Definition

```java
package com.grafana.foundation.dashboardv2beta1.ThresholdsMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


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
