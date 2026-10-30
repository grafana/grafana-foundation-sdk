// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum RecordingRuleHealth {
    UNKNOWN("Unknown"),
    RECORDING("Recording"),
    PAUSED("Paused"),
    ERROR("Error"),
    NO_DATA("NoData"),
    _EMPTY("");

    private final String value;

    private RecordingRuleHealth(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
