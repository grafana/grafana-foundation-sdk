---
title: <span class="badge object-type-enum"></span> TimelineValueAlignment
---
# <span class="badge object-type-enum"></span> TimelineValueAlignment

Controls the value alignment in the TimelineChart component

## Definition

```java
package com.grafana.foundation.common.TimelineValueAlignment;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Controls the value alignment in the TimelineChart component
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum TimelineValueAlignment {
    CENTER("center"),
    LEFT("left"),
    RIGHT("right"),
    _EMPTY("");

    private final String value;

    private TimelineValueAlignment(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
