---
title: <span class="badge object-type-enum"></span> QueryVariableSpecStaticOptionsOrder
---
# <span class="badge object-type-enum"></span> QueryVariableSpecStaticOptionsOrder

## Definition

```java
package com.grafana.foundation.dashboardv2beta1.QueryVariableSpecStaticOptionsOrder;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum QueryVariableSpecStaticOptionsOrder {
    BEFORE("before"),
    AFTER("after"),
    SORTED("sorted"),
    _EMPTY("");

    private final String value;

    private QueryVariableSpecStaticOptionsOrder(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
