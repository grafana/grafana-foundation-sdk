---
title: <span class="badge object-type-enum"></span> GraphTresholdsStyleMode
---
# <span class="badge object-type-enum"></span> GraphTresholdsStyleMode

TODO docs

## Definition

```java
package com.grafana.foundation.common.GraphTresholdsStyleMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum GraphTresholdsStyleMode {
    OFF("off"),
    LINE("line"),
    DASHED("dashed"),
    AREA("area"),
    LINE_AND_AREA("line+area"),
    DASHED_AND_AREA("dashed+area"),
    SERIES("series"),
    _EMPTY("");

    private final String value;

    private GraphTresholdsStyleMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
