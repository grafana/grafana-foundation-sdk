// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum NodesQueryType {
    RANDOM("random"),
    RANDOM_EDGES("random edges"),
    RESPONSE_MEDIUM("response_medium"),
    RESPONSE_SMALL("response_small"),
    FEATURE_SHOWCASE("feature_showcase"),
    _EMPTY("");

    private final String value;

    private NodesQueryType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
