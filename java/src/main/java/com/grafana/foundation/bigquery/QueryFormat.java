// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bigquery;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum QueryFormat {
    TIMESERIES(0),
    TABLE(1);

    private final Integer value;

    private QueryFormat(Integer value) {
        this.value = value;
    }

    @JsonValue
    public Integer Value() {
        return value;
    }
}
