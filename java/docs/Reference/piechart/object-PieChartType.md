---
title: <span class="badge object-type-enum"></span> PieChartType
---
# <span class="badge object-type-enum"></span> PieChartType

Select the pie chart display style.

## Definition

```java
package com.grafana.foundation.piechart.PieChartType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Select the pie chart display style.
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum PieChartType {
    PIE("pie"),
    DONUT("donut"),
    _EMPTY("");

    private final String value;

    private PieChartType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
