---
title: <span class="badge object-type-enum"></span> ColorStrategy
---
# <span class="badge object-type-enum"></span> ColorStrategy

## Definition

```java
package com.grafana.foundation.candlestick.ColorStrategy;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ColorStrategy {
    OPEN_CLOSE("open-close"),
    CLOSE_CLOSE("close-close"),
    _EMPTY("");

    private final String value;

    private ColorStrategy(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
