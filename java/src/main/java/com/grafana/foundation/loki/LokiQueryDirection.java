// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.loki;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum LokiQueryDirection {
    FORWARD("forward"),
    BACKWARD("backward"),
    SCAN("scan"),
    _EMPTY("");

    private final String value;

    private LokiQueryDirection(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
