---
title: <span class="badge object-type-enum"></span> HeatmapColorScale
---
# <span class="badge object-type-enum"></span> HeatmapColorScale

Controls the color scale of the heatmap

## Definition

```java
package com.grafana.foundation.heatmap.HeatmapColorScale;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Controls the color scale of the heatmap
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum HeatmapColorScale {
    LINEAR("linear"),
    EXPONENTIAL("exponential"),
    _EMPTY("");

    private final String value;

    private HeatmapColorScale(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
