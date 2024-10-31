// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum GraphDrawStyle {
    LINE("line"),
    BARS("bars"),
    POINTS("points"),
    _EMPTY("");

    private final String value;

    private GraphDrawStyle(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
