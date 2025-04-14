// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum BuilderQueryEditorPropertyType {
    NUMBER("number"),
    STRING("string"),
    BOOLEAN("boolean"),
    DATETIME("datetime"),
    TIME_SPAN("time_span"),
    FUNCTION("function"),
    INTERVAL("interval"),
    _EMPTY("");

    private final String value;

    private BuilderQueryEditorPropertyType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
