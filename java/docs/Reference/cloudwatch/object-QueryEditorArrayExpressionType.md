---
title: <span class="badge object-type-enum"></span> QueryEditorArrayExpressionType
---
# <span class="badge object-type-enum"></span> QueryEditorArrayExpressionType

## Definition

```java
package com.grafana.foundation.cloudwatch.QueryEditorArrayExpressionType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum QueryEditorArrayExpressionType {
    AND("and"),
    OR("or"),
    _EMPTY("");

    private final String value;

    private QueryEditorArrayExpressionType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
