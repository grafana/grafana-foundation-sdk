// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum TimeSettingsSpecWeekStart {
    SATURDAY("saturday"),
    MONDAY("monday"),
    SUNDAY("sunday"),
    _EMPTY("");

    private final String value;

    private TimeSettingsSpecWeekStart(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
