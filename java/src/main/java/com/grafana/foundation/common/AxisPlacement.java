// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum AxisPlacement {
    AUTO("auto"),
    TOP("top"),
    RIGHT("right"),
    BOTTOM("bottom"),
    LEFT("left"),
    HIDDEN("hidden"),
    _EMPTY("");

    private final String value;

    private AxisPlacement(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
