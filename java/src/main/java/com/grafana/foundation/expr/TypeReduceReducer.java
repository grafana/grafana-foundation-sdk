// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum TypeReduceReducer {
    SUM("sum"),
    MEAN("mean"),
    MIN("min"),
    MAX("max"),
    COUNT("count"),
    LAST("last"),
    MEDIAN("median"),
    _EMPTY("");

    private final String value;

    private TypeReduceReducer(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
