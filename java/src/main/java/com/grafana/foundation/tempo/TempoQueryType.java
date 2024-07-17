// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.tempo;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// search = Loki search, nativeSearch = Tempo search for backwards compatibility
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum TempoQueryType {
    TRACEQL("traceql"),
    TRACEQL_SEARCH("traceqlSearch"),
    SEARCH("search"),
    SERVICE_MAP("serviceMap"),
    UPLOAD("upload"),
    NATIVE_SEARCH("nativeSearch"),
    TRACE_ID("traceId"),
    CLEAR("clear"),
    _EMPTY("");

    private final String value;

    private TempoQueryType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
