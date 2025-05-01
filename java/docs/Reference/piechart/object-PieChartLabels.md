---
title: <span class="badge object-type-enum"></span> PieChartLabels
---
# <span class="badge object-type-enum"></span> PieChartLabels

Select labels to display on the pie chart.

 - Name - The series or field name.

 - Percent - The percentage of the whole.

 - Value - The raw numerical value.

## Definition

```java
package com.grafana.foundation.piechart.PieChartLabels;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Select labels to display on the pie chart.
//  - Name - The series or field name.
//  - Percent - The percentage of the whole.
//  - Value - The raw numerical value.
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum PieChartLabels {
    NAME("name"),
    VALUE("value"),
    PERCENT("percent"),
    _EMPTY("");

    private final String value;

    private PieChartLabels(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
