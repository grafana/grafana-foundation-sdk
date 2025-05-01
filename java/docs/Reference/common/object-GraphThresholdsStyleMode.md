---
title: <span class="badge object-type-enum"></span> GraphThresholdsStyleMode
---
# <span class="badge object-type-enum"></span> GraphThresholdsStyleMode

TODO docs

## Definition

```java
package com.grafana.foundation.common.GraphThresholdsStyleMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum GraphThresholdsStyleMode {
    OFF("off"),
    LINE("line"),
    DASHED("dashed"),
    AREA("area"),
    LINE_AND_AREA("line+area"),
    DASHED_AND_AREA("dashed+area"),
    SERIES("series"),
    _EMPTY("");

    private final String value;

    private GraphThresholdsStyleMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
