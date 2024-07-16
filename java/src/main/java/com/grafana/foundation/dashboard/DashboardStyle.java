// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum DashboardStyle {
    LIGHT("light"),
    DARK("dark"),
    _EMPTY("");

    private final String value;

    private DashboardStyle(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
