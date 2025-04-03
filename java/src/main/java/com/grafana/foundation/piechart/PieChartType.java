// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.piechart;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Select the pie chart display style.
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum PieChartType {
    PIE("pie"),
    DONUT("donut"),
    _EMPTY("");

    private final String value;

    private PieChartType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
