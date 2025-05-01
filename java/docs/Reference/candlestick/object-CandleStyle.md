---
title: <span class="badge object-type-enum"></span> CandleStyle
---
# <span class="badge object-type-enum"></span> CandleStyle

## Definition

```java
package com.grafana.foundation.candlestick.CandleStyle;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum CandleStyle {
    CANDLES("candles"),
    OHLC_BARS("ohlcbars"),
    _EMPTY("");

    private final String value;

    private CandleStyle(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
