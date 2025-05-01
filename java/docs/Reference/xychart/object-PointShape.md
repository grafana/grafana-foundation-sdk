---
title: <span class="badge object-type-enum"></span> PointShape
---
# <span class="badge object-type-enum"></span> PointShape

## Definition

```java
package com.grafana.foundation.xychart.PointShape;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum PointShape {
    CIRCLE("circle"),
    SQUARE("square"),
    _EMPTY("");

    private final String value;

    private PointShape(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
