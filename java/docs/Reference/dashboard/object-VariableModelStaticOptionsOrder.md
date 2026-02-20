---
title: <span class="badge object-type-enum"></span> VariableModelStaticOptionsOrder
---
# <span class="badge object-type-enum"></span> VariableModelStaticOptionsOrder

## Definition

```java
package com.grafana.foundation.dashboard.VariableModelStaticOptionsOrder;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum VariableModelStaticOptionsOrder {
    BEFORE("before"),
    AFTER("after"),
    SORTED("sorted"),
    _EMPTY("");

    private final String value;

    private VariableModelStaticOptionsOrder(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
