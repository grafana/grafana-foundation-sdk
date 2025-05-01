---
title: <span class="badge object-type-enum"></span> DashboardStyle
---
# <span class="badge object-type-enum"></span> DashboardStyle

## Definition

```java
package com.grafana.foundation.dashboard.DashboardStyle;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum DashboardStyle {
    LIGHT("light"),
    DARK("dark"),
    _EMPTY("");

    private final String value;

    private DashboardStyle(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
