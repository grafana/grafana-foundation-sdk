// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum HorizontalConstraint {
    LEFT("left"),
    RIGHT("right"),
    LEFT_RIGHT("leftright"),
    CENTER("center"),
    SCALE("scale"),
    _EMPTY("");

    private final String value;

    private HorizontalConstraint(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
