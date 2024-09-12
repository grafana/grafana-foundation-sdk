// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.geomap;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum TooltipMode {
    NONE("none"),
    DETAILS("details"),
    _EMPTY("");

    private final String value;

    private TooltipMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
