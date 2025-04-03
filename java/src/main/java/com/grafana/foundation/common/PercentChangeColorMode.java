// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum PercentChangeColorMode {
    STANDARD("standard"),
    INVERTED("inverted"),
    SAME_AS_VALUE("same_as_value"),
    _EMPTY("");

    private final String value;

    private PercentChangeColorMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
