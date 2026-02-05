---
title: <span class="badge object-type-enum"></span> QueryEditorExpressionType
---
# <span class="badge object-type-enum"></span> QueryEditorExpressionType

## Definition

```java
package com.grafana.foundation.bigquery.QueryEditorExpressionType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum QueryEditorExpressionType {
    PROPERTY("property"),
    OPERATOR("operator"),
    OR("or"),
    AND("and"),
    GROUP_BY("groupBy"),
    FUNCTION("function"),
    FUNCTION_PARAMETER("functionParameter"),
    _EMPTY("");

    private final String value;

    private QueryEditorExpressionType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
