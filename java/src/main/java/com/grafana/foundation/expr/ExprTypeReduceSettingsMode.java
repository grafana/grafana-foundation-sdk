// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ExprTypeReduceSettingsMode {
    DROP_NN("dropNN"),
    REPLACE_NN("replaceNN"),
    _EMPTY("");

    private final String value;

    private ExprTypeReduceSettingsMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
