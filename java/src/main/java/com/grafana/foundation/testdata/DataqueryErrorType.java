// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum DataqueryErrorType {
    SERVER_PANIC("server_panic"),
    FRONTEND_EXCEPTION("frontend_exception"),
    FRONTEND_OBSERVABLE("frontend_observable"),
    _EMPTY("");

    private final String value;

    private DataqueryErrorType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
