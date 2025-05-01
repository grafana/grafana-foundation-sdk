---
title: <span class="badge object-type-enum"></span> PromQueryFormat
---
# <span class="badge object-type-enum"></span> PromQueryFormat

## Definition

```java
package com.grafana.foundation.prometheus.PromQueryFormat;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum PromQueryFormat {
    TIME_SERIES("time_series"),
    TABLE("table"),
    HEATMAP("heatmap"),
    _EMPTY("");

    private final String value;

    private PromQueryFormat(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
