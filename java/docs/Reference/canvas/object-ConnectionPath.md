---
title: <span class="badge object-type-enum"></span> ConnectionPath
---
# <span class="badge object-type-enum"></span> ConnectionPath

## Definition

```java
package com.grafana.foundation.canvas.ConnectionPath;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ConnectionPath {
    STRAIGHT("straight"),
    _EMPTY("");

    private final String value;

    private ConnectionPath(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
