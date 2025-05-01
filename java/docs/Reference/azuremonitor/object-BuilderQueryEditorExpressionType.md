---
title: <span class="badge object-type-enum"></span> BuilderQueryEditorExpressionType
---
# <span class="badge object-type-enum"></span> BuilderQueryEditorExpressionType

## Definition

```java
package com.grafana.foundation.azuremonitor.BuilderQueryEditorExpressionType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum BuilderQueryEditorExpressionType {
    PROPERTY("property"),
    OPERATOR("operator"),
    REDUCE("reduce"),
    FUNCTION_PARAMETER("function_parameter"),
    GROUP_BY("group_by"),
    OR("or"),
    AND("and"),
    ORDER_BY("order_by"),
    _EMPTY("");

    private final String value;

    private BuilderQueryEditorExpressionType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
