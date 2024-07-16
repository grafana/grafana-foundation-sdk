// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.candlestick;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ColorStrategy {
    OPEN_CLOSE("open-close"),
    CLOSE_CLOSE("close-close"),
    _EMPTY("");

    private final String value;

    private ColorStrategy(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
