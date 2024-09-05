// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum TypeClassicConditionsType {
    AND("and"),
    OR("or"),
    LOGIC_OR("logic-or"),
    _EMPTY("");

    private final String value;

    private TypeClassicConditionsType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
