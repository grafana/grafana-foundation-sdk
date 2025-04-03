// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Allows for the bar gauge name to be placed explicitly
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum BarGaugeNamePlacement {
    AUTO("auto"),
    TOP("top"),
    LEFT("left"),
    HIDDEN("hidden"),
    _EMPTY("");

    private final String value;

    private BarGaugeNamePlacement(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
