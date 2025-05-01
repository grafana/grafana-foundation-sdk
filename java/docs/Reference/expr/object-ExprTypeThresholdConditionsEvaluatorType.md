---
title: <span class="badge object-type-enum"></span> ExprTypeThresholdConditionsEvaluatorType
---
# <span class="badge object-type-enum"></span> ExprTypeThresholdConditionsEvaluatorType

## Definition

```java
package com.grafana.foundation.expr.ExprTypeThresholdConditionsEvaluatorType;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ExprTypeThresholdConditionsEvaluatorType {
    GT("gt"),
    LT("lt"),
    EQ("eq"),
    NE("ne"),
    GTE("gte"),
    LTE("lte"),
    WITHIN_RANGE("within_range"),
    OUTSIDE_RANGE("outside_range"),
    WITHIN_RANGE_INCLUDED("within_range_included"),
    OUTSIDE_RANGE_INCLUDED("outside_range_included"),
    _EMPTY("");

    private final String value;

    private ExprTypeThresholdConditionsEvaluatorType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}

```
