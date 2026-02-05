---
title: <span class="badge object-type-enum"></span> RepeatOptionsDirection
---
# <span class="badge object-type-enum"></span> RepeatOptionsDirection

## Definition

```java
package com.grafana.foundation.dashboardv2beta1.RepeatOptionsDirection;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum RepeatOptionsDirection {
    H("h"),
    V("v"),
    _EMPTY("");

    private final String value;

    private RepeatOptionsDirection(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
