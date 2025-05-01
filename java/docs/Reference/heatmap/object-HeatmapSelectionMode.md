---
title: <span class="badge object-type-enum"></span> HeatmapSelectionMode
---
# <span class="badge object-type-enum"></span> HeatmapSelectionMode

Controls which axis to allow selection on

## Definition

```java
package com.grafana.foundation.heatmap.HeatmapSelectionMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Controls which axis to allow selection on
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum HeatmapSelectionMode {
    X("x"),
    Y("y"),
    XY("xy"),
    _EMPTY("");

    private final String value;

    private HeatmapSelectionMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
