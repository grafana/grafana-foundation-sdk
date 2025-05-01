---
title: <span class="badge object-type-enum"></span> ScalarDimensionMode
---
# <span class="badge object-type-enum"></span> ScalarDimensionMode

## Definition

```java
package com.grafana.foundation.common.ScalarDimensionMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ScalarDimensionMode {
    MOD("mod"),
    CLAMPED("clamped"),
    _EMPTY("");

    private final String value;

    private ScalarDimensionMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
