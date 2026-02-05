// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum DataqueryErrorSource {
    PLUGIN("plugin"),
    DOWNSTREAM("downstream"),
    _EMPTY("");

    private final String value;

    private DataqueryErrorSource(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
