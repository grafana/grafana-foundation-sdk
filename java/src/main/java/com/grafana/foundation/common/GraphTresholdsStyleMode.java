// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum GraphTresholdsStyleMode {
    OFF("off"),
    LINE("line"),
    DASHED("dashed"),
    AREA("area"),
    LINE_AND_AREA("line+area"),
    DASHED_AND_AREA("dashed+area"),
    SERIES("series"),
    _EMPTY("");

    private final String value;

    private GraphTresholdsStyleMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
