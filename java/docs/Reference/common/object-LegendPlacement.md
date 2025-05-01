---
title: <span class="badge object-type-enum"></span> LegendPlacement
---
# <span class="badge object-type-enum"></span> LegendPlacement

TODO docs

## Definition

```java
package com.grafana.foundation.common.LegendPlacement;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum LegendPlacement {
    BOTTOM("bottom"),
    RIGHT("right"),
    _EMPTY("");

    private final String value;

    private LegendPlacement(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
