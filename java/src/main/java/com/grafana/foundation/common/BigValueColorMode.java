// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum BigValueColorMode {
    VALUE("value"),
    BACKGROUND("background"),
    BACKGROUND_SOLID("background_solid"),
    NONE("none"),
    _EMPTY("");

    private final String value;

    private BigValueColorMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
