---
title: <span class="badge object-type-enum"></span> LogsEditorMode
---
# <span class="badge object-type-enum"></span> LogsEditorMode

## Definition

```java
package com.grafana.foundation.azuremonitor.LogsEditorMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum LogsEditorMode {
    BUILDER("builder"),
    RAW("raw"),
    _EMPTY("");

    private final String value;

    private LogsEditorMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
