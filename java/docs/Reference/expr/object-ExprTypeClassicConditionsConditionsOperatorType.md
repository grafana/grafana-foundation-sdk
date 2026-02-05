---
title: <span class="badge object-type-enum"></span> ExprTypeClassicConditionsConditionsOperatorType
---
# <span class="badge object-type-enum"></span> ExprTypeClassicConditionsConditionsOperatorType

## Definition

```java
package com.grafana.foundation.expr.ExprTypeClassicConditionsConditionsOperatorType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ExprTypeClassicConditionsConditionsOperatorType {
    AND("and"),
    OR("or"),
    LOGIC_OR("logic-or"),
    _EMPTY("");

    private final String value;

    private ExprTypeClassicConditionsConditionsOperatorType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
