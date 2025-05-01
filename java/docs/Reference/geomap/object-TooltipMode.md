---
title: <span class="badge object-type-enum"></span> TooltipMode
---
# <span class="badge object-type-enum"></span> TooltipMode

## Definition

```java
package com.grafana.foundation.geomap.TooltipMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum TooltipMode {
    NONE("none"),
    DETAILS("details"),
    _EMPTY("");

    private final String value;

    private TooltipMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
