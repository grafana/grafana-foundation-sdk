---
title: <span class="badge object-type-enum"></span> VizDisplayMode
---
# <span class="badge object-type-enum"></span> VizDisplayMode

## Definition

```java
package com.grafana.foundation.candlestick.VizDisplayMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum VizDisplayMode {
    CANDLES_VOLUME("candles+volume"),
    CANDLES("candles"),
    VOLUME("volume"),
    _EMPTY("");

    private final String value;

    private VizDisplayMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
