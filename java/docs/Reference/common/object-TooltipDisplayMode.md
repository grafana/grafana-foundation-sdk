---
title: <span class="badge object-type-enum"></span> TooltipDisplayMode
---
# <span class="badge object-type-enum"></span> TooltipDisplayMode

TODO docs

## Definition

```java
package com.grafana.foundation.common.TooltipDisplayMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum TooltipDisplayMode {
    SINGLE("single"),
    MULTI("multi"),
    NONE("none"),
    _EMPTY("");

    private final String value;

    private TooltipDisplayMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
