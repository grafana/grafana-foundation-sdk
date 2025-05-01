---
title: <span class="badge object-type-enum"></span> AxisPlacement
---
# <span class="badge object-type-enum"></span> AxisPlacement

TODO docs

## Definition

```java
package com.grafana.foundation.common.AxisPlacement;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum AxisPlacement {
    AUTO("auto"),
    TOP("top"),
    RIGHT("right"),
    BOTTOM("bottom"),
    LEFT("left"),
    HIDDEN("hidden"),
    _EMPTY("");

    private final String value;

    private AxisPlacement(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
