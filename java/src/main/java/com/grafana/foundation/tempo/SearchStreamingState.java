// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.tempo;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// The state of the TraceQL streaming search query
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum SearchStreamingState {
    PENDING("pending"),
    STREAMING("streaming"),
    DONE("done"),
    ERROR("error"),
    _EMPTY("");

    private final String value;

    private SearchStreamingState(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
