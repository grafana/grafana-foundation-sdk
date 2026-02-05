---
title: <span class="badge object-type-enum"></span> HeatmapCalculationMode
---
# <span class="badge object-type-enum"></span> HeatmapCalculationMode

## Definition

```java
package com.grafana.foundation.common.HeatmapCalculationMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum HeatmapCalculationMode {
    SIZE("size"),
    COUNT("count"),
    _EMPTY("");

    private final String value;

    private HeatmapCalculationMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
