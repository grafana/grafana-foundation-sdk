// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum RuleExecErrState {
    OK("OK"),
    ALERTING("Alerting"),
    ERROR("Error"),
    _EMPTY("");

    private final String value;

    private RuleExecErrState(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
