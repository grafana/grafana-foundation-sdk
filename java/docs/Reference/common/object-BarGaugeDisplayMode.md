---
title: <span class="badge object-type-enum"></span> BarGaugeDisplayMode
---
# <span class="badge object-type-enum"></span> BarGaugeDisplayMode

Enum expressing the possible display modes

for the bar gauge component of Grafana UI

## Definition

```java
package com.grafana.foundation.common.BarGaugeDisplayMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Enum expressing the possible display modes
// for the bar gauge component of Grafana UI
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum BarGaugeDisplayMode {
    BASIC("basic"),
    LCD("lcd"),
    GRADIENT("gradient"),
    _EMPTY("");

    private final String value;

    private BarGaugeDisplayMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
