// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Compare two values
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ComparisonOperation {
    EQ("eq"),
    NEQ("neq"),
    LT("lt"),
    LTE("lte"),
    GT("gt"),
    GTE("gte"),
    _EMPTY("");

    private final String value;

    private ComparisonOperation(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
