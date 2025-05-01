---
title: <span class="badge object-type-enum"></span> BarGaugeNamePlacement
---
# <span class="badge object-type-enum"></span> BarGaugeNamePlacement

Allows for the bar gauge name to be placed explicitly

## Definition

```java
package com.grafana.foundation.common.BarGaugeNamePlacement;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Allows for the bar gauge name to be placed explicitly
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum BarGaugeNamePlacement {
    AUTO("auto"),
    TOP("top"),
    LEFT("left"),
    _EMPTY("");

    private final String value;

    private BarGaugeNamePlacement(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
