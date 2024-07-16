// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum MovingAverageModel {
    SIMPLE("simple"),
    LINEAR("linear"),
    EWMA("ewma"),
    HOLT("holt"),
    HOLT_WINTERS("holt_winters"),
    _EMPTY("");

    private final String value;

    private MovingAverageModel(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
