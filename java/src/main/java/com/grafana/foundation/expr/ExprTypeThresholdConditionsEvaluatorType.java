// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ExprTypeThresholdConditionsEvaluatorType {
    GT("gt"),
    LT("lt"),
    WITHIN_RANGE("within_range"),
    OUTSIDE_RANGE("outside_range"),
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
