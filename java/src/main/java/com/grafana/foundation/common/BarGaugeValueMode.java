// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Allows for the table cell gauge display type to set the gauge mode.
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum BarGaugeValueMode {
    COLOR("color"),
    TEXT("text"),
    HIDDEN("hidden"),
    _EMPTY("");

    private final String value;

    private BarGaugeValueMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
