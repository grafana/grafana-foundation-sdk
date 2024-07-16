// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum TypeReduceType {
    NONE(""),
    TIMESERIES_WIDE("timeseries-wide"),
    TIMESERIES_LONG("timeseries-long"),
    TIMESERIES_MANY("timeseries-many"),
    TIMESERIES_MULTI("timeseries-multi"),
    DIRECTORY_LISTING("directory-listing"),
    TABLE("table"),
    NUMERIC_WIDE("numeric-wide"),
    NUMERIC_MULTI("numeric-multi"),
    NUMERIC_LONG("numeric-long"),
    LOG_LINES("log-lines");

    private final String value;

    private TypeReduceType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
