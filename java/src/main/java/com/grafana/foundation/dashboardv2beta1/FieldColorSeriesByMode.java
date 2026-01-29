// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Defines how to assign a series color from "by value" color schemes. For example for an aggregated data points like a timeseries, the color can be assigned by the min, max or last value.
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum FieldColorSeriesByMode {
    MIN("min"),
    MAX("max"),
    LAST("last"),
    _EMPTY("");

    private final String value;

    private FieldColorSeriesByMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
