---
title: <span class="badge object-type-enum"></span> PanelModelRepeatDirection
---
# <span class="badge object-type-enum"></span> PanelModelRepeatDirection

## Definition

```java
package com.grafana.foundation.librarypanel.PanelModelRepeatDirection;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum PanelModelRepeatDirection {
    H("h"),
    V("v"),
    _EMPTY("");

    private final String value;

    private PanelModelRepeatDirection(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
