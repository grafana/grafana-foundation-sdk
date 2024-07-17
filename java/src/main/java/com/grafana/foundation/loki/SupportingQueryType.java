// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.loki;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum SupportingQueryType {
    LOGS_VOLUME("logsVolume"),
    LOGS_SAMPLE("logsSample"),
    DATA_SAMPLE("dataSample"),
    INFINITE_SCROLL("infiniteScroll"),
    _EMPTY("");

    private final String value;

    private SupportingQueryType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
