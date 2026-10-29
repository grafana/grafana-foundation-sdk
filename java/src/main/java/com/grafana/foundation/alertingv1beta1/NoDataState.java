// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum NoDataState {
    NO_DATA("NoData"),
    OK("Ok"),
    ALERTING("Alerting"),
    KEEP_LAST("KeepLast"),
    _EMPTY("");

    private final String value;

    private NoDataState(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
