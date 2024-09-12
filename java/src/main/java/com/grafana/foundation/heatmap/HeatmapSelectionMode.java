// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Controls which axis to allow selection on
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum HeatmapSelectionMode {
    X("x"),
    Y("y"),
    XY("xy"),
    _EMPTY("");

    private final String value;

    private HeatmapSelectionMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
