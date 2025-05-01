---
title: <span class="badge object-type-enum"></span> LineStyleFill
---
# <span class="badge object-type-enum"></span> LineStyleFill

## Definition

```java
package com.grafana.foundation.common.LineStyleFill;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum LineStyleFill {
    SOLID("solid"),
    DASH("dash"),
    DOT("dot"),
    SQUARE("square"),
    _EMPTY("");

    private final String value;

    private LineStyleFill(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
