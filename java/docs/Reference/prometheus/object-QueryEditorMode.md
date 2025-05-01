---
title: <span class="badge object-type-enum"></span> QueryEditorMode
---
# <span class="badge object-type-enum"></span> QueryEditorMode

## Definition

```java
package com.grafana.foundation.prometheus.QueryEditorMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum QueryEditorMode {
    CODE("code"),
    BUILDER("builder"),
    _EMPTY("");

    private final String value;

    private QueryEditorMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
