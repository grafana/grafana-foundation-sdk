// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Defines the supported queryTypes.
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum QueryType {
    TIME_SERIES_LIST("timeSeriesList"),
    TIME_SERIES_QUERY("timeSeriesQuery"),
    SLO("slo"),
    ANNOTATION("annotation"),
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
