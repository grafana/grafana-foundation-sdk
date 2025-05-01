---
title: <span class="badge object-type-enum"></span> BuilderQueryEditorPropertyType
---
# <span class="badge object-type-enum"></span> BuilderQueryEditorPropertyType

## Definition

```java
package com.grafana.foundation.azuremonitor.BuilderQueryEditorPropertyType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum BuilderQueryEditorPropertyType {
    NUMBER("number"),
    STRING("string"),
    BOOLEAN("boolean"),
    DATETIME("datetime"),
    TIME_SPAN("time_span"),
    FUNCTION("function"),
    INTERVAL("interval"),
    _EMPTY("");

    private final String value;

    private BuilderQueryEditorPropertyType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
