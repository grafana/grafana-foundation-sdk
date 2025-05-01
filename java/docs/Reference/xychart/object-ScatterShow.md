---
title: <span class="badge object-type-enum"></span> ScatterShow
---
# <span class="badge object-type-enum"></span> ScatterShow

## Definition

```java
package com.grafana.foundation.xychart.ScatterShow;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ScatterShow {
    POINTS("points"),
    LINES("lines"),
    POINTS_AND_LINES("points+lines"),
    _EMPTY("");

    private final String value;

    private ScatterShow(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
