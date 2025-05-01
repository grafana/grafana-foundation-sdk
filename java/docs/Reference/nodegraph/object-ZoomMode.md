---
title: <span class="badge object-type-enum"></span> ZoomMode
---
# <span class="badge object-type-enum"></span> ZoomMode

## Definition

```java
package com.grafana.foundation.nodegraph.ZoomMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ZoomMode {
    COOPERATIVE("cooperative"),
    GREEDY("greedy"),
    _EMPTY("");

    private final String value;

    private ZoomMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
