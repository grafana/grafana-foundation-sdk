// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Controls the color mode of the heatmap
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum HeatmapColorMode {
    OPACITY("opacity"),
    SCHEME("scheme"),
    _EMPTY("");

    private final String value;

    private HeatmapColorMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
