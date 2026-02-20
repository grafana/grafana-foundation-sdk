---
title: <span class="badge object-type-enum"></span> ActionVariableType
---
# <span class="badge object-type-enum"></span> ActionVariableType

Action variable type

## Definition

```java
package com.grafana.foundation.dashboard.ActionVariableType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Action variable type
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ActionVariableType {
    STRING("string"),
    _EMPTY("");

    private final String value;

    private ActionVariableType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
