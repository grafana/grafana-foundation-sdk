---
title: <span class="badge object-type-enum"></span> TextDimensionMode
---
# <span class="badge object-type-enum"></span> TextDimensionMode

## Definition

```java
package com.grafana.foundation.common.TextDimensionMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum TextDimensionMode {
    FIXED("fixed"),
    FIELD("field"),
    TEMPLATE("template"),
    _EMPTY("");

    private final String value;

    private TextDimensionMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
