---
title: <span class="badge object-type-enum"></span> XYShowMode
---
# <span class="badge object-type-enum"></span> XYShowMode

## Definition

```java
package com.grafana.foundation.xychart.XYShowMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum XYShowMode {
    POINTS("points"),
    LINES("lines"),
    POINTS_AND_LINES("points+lines"),
    _EMPTY("");

    private final String value;

    private XYShowMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
