---
title: <span class="badge object-type-enum"></span> OptionsDetailsMode
---
# <span class="badge object-type-enum"></span> OptionsDetailsMode

## Definition

```java
package com.grafana.foundation.logs.OptionsDetailsMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum OptionsDetailsMode {
    INLINE("inline"),
    SIDEBAR("sidebar"),
    _EMPTY("");

    private final String value;

    private OptionsDetailsMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
