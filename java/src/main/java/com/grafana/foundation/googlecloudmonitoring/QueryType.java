// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Defines the supported queryTypes.
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum QueryType {
    TIMESERIESLIST("timeSeriesList"),
    TIMESERIESQUERY("timeSeriesQuery"),
    SLO("slo"),
    ANNOTATION("annotation"),
    PROMQL("promQL"),
    _EMPTY("");

    private final String value;

    private QueryType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
