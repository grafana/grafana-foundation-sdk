// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Enum expressing the possible display modes
// for the bar gauge component of Grafana UI
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum BarGaugeDisplayMode {
    BASIC("basic"),
    LCD("lcd"),
    GRADIENT("gradient"),
    _EMPTY("");

    private final String value;

    private BarGaugeDisplayMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
