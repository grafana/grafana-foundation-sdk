// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum FrameGeometrySourceMode {
    AUTO("auto"),
    GEOHASH("geohash"),
    COORDS("coords"),
    LOOKUP("lookup"),
    _EMPTY("");

    private final String value;

    private FrameGeometrySourceMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
