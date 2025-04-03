// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// TODO docs
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ScaleDirection {
    UP(1),
    RIGHT(1),
    DOWN(-1),
    LEFT(-1);

    private final Integer value;

    private ScaleDirection(Integer value) {
        this.value = value;
    }

    @JsonValue
    public Integer Value() {
        return value;
    }
}
