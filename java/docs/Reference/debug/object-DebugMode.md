---
title: <span class="badge object-type-enum"></span> DebugMode
---
# <span class="badge object-type-enum"></span> DebugMode

## Definition

```java
package com.grafana.foundation.debug.DebugMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum DebugMode {
    RENDER("render"),
    EVENTS("events"),
    CURSOR("cursor"),
    STATE("State"),
    THROW_ERROR("ThrowError"),
    _EMPTY("");

    private final String value;

    private DebugMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
