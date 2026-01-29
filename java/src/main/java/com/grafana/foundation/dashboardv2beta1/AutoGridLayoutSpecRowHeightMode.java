// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum AutoGridLayoutSpecRowHeightMode {
    SHORT("short"),
    STANDARD("standard"),
    TALL("tall"),
    CUSTOM("custom"),
    _EMPTY("");

    private final String value;

    private AutoGridLayoutSpecRowHeightMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
