---
title: <span class="badge object-type-enum"></span> ActionType
---
# <span class="badge object-type-enum"></span> ActionType

Dashboard action type

## Definition

```java
package com.grafana.foundation.dashboard.ActionType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Dashboard action type
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ActionType {
    FETCH("fetch"),
    INFINITY("infinity"),
    _EMPTY("");

    private final String value;

    private ActionType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
