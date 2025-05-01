---
title: <span class="badge object-type-enum"></span> BackgroundImageSize
---
# <span class="badge object-type-enum"></span> BackgroundImageSize

## Definition

```java
package com.grafana.foundation.canvas.BackgroundImageSize;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum BackgroundImageSize {
    ORIGINAL("original"),
    CONTAIN("contain"),
    COVER("cover"),
    FILL("fill"),
    TILE("tile"),
    _EMPTY("");

    private final String value;

    private BackgroundImageSize(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
