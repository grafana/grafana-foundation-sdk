// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.parca;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ParcaQueryType {
    METRICS("metrics"),
    PROFILE("profile"),
    BOTH("both"),
    _EMPTY("");

    private final String value;

    private ParcaQueryType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
