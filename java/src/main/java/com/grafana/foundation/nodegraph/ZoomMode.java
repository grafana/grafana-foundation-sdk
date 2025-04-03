// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.nodegraph;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ZoomMode {
    COOPERATIVE("cooperative"),
    GREEDY("greedy"),
    _EMPTY("");

    private final String value;

    private ZoomMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
