// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum HttpRequestMethod {
    GET("GET"),
    POST("POST"),
    PUT("PUT"),
    _EMPTY("");

    private final String value;

    private HttpRequestMethod(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
