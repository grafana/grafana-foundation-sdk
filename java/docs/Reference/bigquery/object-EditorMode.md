---
title: <span class="badge object-type-enum"></span> EditorMode
---
# <span class="badge object-type-enum"></span> EditorMode

## Definition

```java
package com.grafana.foundation.bigquery.EditorMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum EditorMode {
    CODE("code"),
    BUILDER("builder"),
    _EMPTY("");

    private final String value;

    private EditorMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
