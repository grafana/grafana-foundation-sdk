// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Loading status
// Accepted values are `NotStarted` (the request is not started), `Loading` (waiting for response), `Streaming` (pulling continuous data), `Done` (response received successfully) or `Error` (failed request).
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum LoadingState {
    NOT_STARTED("NotStarted"),
    LOADING("Loading"),
    STREAMING("Streaming"),
    DONE("Done"),
    ERROR("Error"),
    _EMPTY("");

    private final String value;

    private LoadingState(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
