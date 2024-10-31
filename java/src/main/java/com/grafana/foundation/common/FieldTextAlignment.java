// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO -- should not be table specific!
// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum FieldTextAlignment {
    AUTO("auto"),
    LEFT("left"),
    RIGHT("right"),
    CENTER("center"),
    _EMPTY("");

    private final String value;

    private FieldTextAlignment(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
