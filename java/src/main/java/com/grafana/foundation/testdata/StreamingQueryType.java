// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum StreamingQueryType {
    SIGNAL("signal"),
    LOGS("logs"),
    FETCH("fetch"),
    _EMPTY("");

    private final String value;

    private StreamingQueryType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
