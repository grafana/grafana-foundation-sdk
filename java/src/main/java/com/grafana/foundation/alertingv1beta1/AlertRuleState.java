// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum AlertRuleState {
    INACTIVE("Inactive"),
    HEALTHY("Healthy"),
    FIRING("Firing"),
    PENDING("Pending"),
    RECOVERING("Recovering"),
    _EMPTY("");

    private final String value;

    private AlertRuleState(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
