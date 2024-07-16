// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Determine if the variable shows on dashboard
// Accepted values are 0 (show label and value), 1 (show value only), 2 (show nothing).
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum VariableHide {
    DONT_HIDE(0),
    HIDE_LABEL(1),
    HIDE_VARIABLE(2);

    private final Integer value;

    private VariableHide(Integer value) {
        this.value = value;
    }

    @JsonValue
    public Integer Value() {
        return value;
    }
}
