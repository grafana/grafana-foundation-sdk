---
title: <span class="badge object-type-enum"></span> LokiQueryDirection
---
# <span class="badge object-type-enum"></span> LokiQueryDirection

## Definition

```java
package com.grafana.foundation.loki.LokiQueryDirection;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum LokiQueryDirection {
    FORWARD("forward"),
    BACKWARD("backward"),
    _EMPTY("");

    private final String value;

    private LokiQueryDirection(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
