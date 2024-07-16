// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum QueryEditorPropertyType {
    STRING("string"),
    _EMPTY("");

    private final String value;

    private QueryEditorPropertyType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
