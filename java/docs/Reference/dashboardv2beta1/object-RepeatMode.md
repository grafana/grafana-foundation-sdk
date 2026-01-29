---
title: <span class="badge object-type-enum"></span> RepeatMode
---
# <span class="badge object-type-enum"></span> RepeatMode

other repeat modes will be added in the future: label, frame

## Definition

```java
package com.grafana.foundation.dashboardv2beta1.RepeatMode;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// other repeat modes will be added in the future: label, frame
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum RepeatMode {
    VARIABLE("variable"),
    _EMPTY("");

    private final String value;

    private RepeatMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
