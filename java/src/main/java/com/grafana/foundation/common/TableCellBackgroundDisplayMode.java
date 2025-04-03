// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Display mode to the "Colored Background" display
// mode for table cells. Either displays a solid color (basic mode)
// or a gradient.
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum TableCellBackgroundDisplayMode {
    BASIC("basic"),
    GRADIENT("gradient"),
    _EMPTY("");

    private final String value;

    private TableCellBackgroundDisplayMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
