// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Action variable type
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ActionVariableType {
    STRING("string"),
    _EMPTY("");

    private final String value;

    private ActionVariableType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
