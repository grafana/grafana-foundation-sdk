// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum RuleNoDataState {
    OK("OK"),
    ALERTING("Alerting"),
    NO_DATA("NoData"),
    KEEP_LAST("KeepLast"),
    _EMPTY("");

    private final String value;

    private RuleNoDataState(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
