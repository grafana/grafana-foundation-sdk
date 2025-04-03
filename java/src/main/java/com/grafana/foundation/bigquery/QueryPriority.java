// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bigquery;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum QueryPriority {
    INTERACTIVE("INTERACTIVE"),
    BATCH("BATCH"),
    _EMPTY("");

    private final String value;

    private QueryPriority(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
