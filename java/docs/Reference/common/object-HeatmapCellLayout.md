---
title: <span class="badge object-type-enum"></span> HeatmapCellLayout
---
# <span class="badge object-type-enum"></span> HeatmapCellLayout

## Definition

```java
package com.grafana.foundation.common.HeatmapCellLayout;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum HeatmapCellLayout {
    LE("le"),
    GE("ge"),
    UNKNOWN("unknown"),
    AUTO("auto"),
    _EMPTY("");

    private final String value;

    private HeatmapCellLayout(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
