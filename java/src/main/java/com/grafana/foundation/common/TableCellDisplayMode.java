// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Internally, this is the "type" of cell that's being displayed
// in the table such as colored text, JSON, gauge, etc.
// The color-background-solid, gradient-gauge, and lcd-gauge
// modes are deprecated in favor of new cell subOptions
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum TableCellDisplayMode {
    AUTO("auto"),
    COLOR_TEXT("color-text"),
    COLOR_BACKGROUND("color-background"),
    COLOR_BACKGROUND_SOLID("color-background-solid"),
    GRADIENT_GAUGE("gradient-gauge"),
    LCD_GAUGE("lcd-gauge"),
    JSON_VIEW("json-view"),
    BASIC_GAUGE("basic"),
    IMAGE("image"),
    GAUGE("gauge"),
    SPARKLINE("sparkline"),
    DATA_LINKS("data-links"),
    CUSTOM("custom"),
    ACTIONS("actions"),
    _EMPTY("");

    private final String value;

    private TableCellDisplayMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
