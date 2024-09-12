// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.debug;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum DebugMode {
    RENDER("render"),
    EVENTS("events"),
    CURSOR("cursor"),
    STATE("State"),
    THROW_ERROR("ThrowError"),
    _EMPTY("");

    private final String value;

    private DebugMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
