---
title: <span class="badge object-type-enum"></span> HeatmapColorMode
---
# <span class="badge object-type-enum"></span> HeatmapColorMode

Controls the color mode of the heatmap

## Definition

```java
package com.grafana.foundation.heatmap.HeatmapColorMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Controls the color mode of the heatmap
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum HeatmapColorMode {
    OPACITY("opacity"),
    SCHEME("scheme"),
    _EMPTY("");

    private final String value;

    private HeatmapColorMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
