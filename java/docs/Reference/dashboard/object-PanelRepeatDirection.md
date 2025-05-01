---
title: <span class="badge object-type-enum"></span> PanelRepeatDirection
---
# <span class="badge object-type-enum"></span> PanelRepeatDirection

## Definition

```java
package com.grafana.foundation.dashboard.PanelRepeatDirection;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum PanelRepeatDirection {
    H("h"),
    V("v"),
    _EMPTY("");

    private final String value;

    private PanelRepeatDirection(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
