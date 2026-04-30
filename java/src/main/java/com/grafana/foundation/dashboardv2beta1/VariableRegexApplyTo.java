// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Determine whether regex applies to variable value or display text
// Accepted values are `value` (apply to value used in queries) or `text` (apply to display text shown to users)
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum VariableRegexApplyTo {
    VALUE("value"),
    TEXT("text"),
    _EMPTY("");

    private final String value;

    private VariableRegexApplyTo(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
