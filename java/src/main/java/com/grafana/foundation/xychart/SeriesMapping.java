// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Auto is "table" in the UI
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum SeriesMapping {
    AUTO("auto"),
    MANUAL("manual"),
    _EMPTY("");

    private final String value;

    private SeriesMapping(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
