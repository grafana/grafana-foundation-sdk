// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.grafanapyroscope;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum PhlareQueryType {
    METRICS("metrics"),
    PROFILE("profile"),
    BOTH("both"),
    _EMPTY("");

    private final String value;

    private PhlareQueryType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
