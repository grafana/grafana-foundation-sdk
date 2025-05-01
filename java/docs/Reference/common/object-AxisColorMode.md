---
title: <span class="badge object-type-enum"></span> AxisColorMode
---
# <span class="badge object-type-enum"></span> AxisColorMode

TODO docs

## Definition

```java
package com.grafana.foundation.common.AxisColorMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum AxisColorMode {
    TEXT("text"),
    SERIES("series"),
    _EMPTY("");

    private final String value;

    private AxisColorMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
