// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.piechart;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Select values to display in the legend.
//  - Percent: The percentage of the whole.
//  - Value: The raw numerical value.
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum PieChartLegendValues {
    VALUE("value"),
    PERCENT("percent"),
    _EMPTY("");

    private final String value;

    private PieChartLegendValues(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
