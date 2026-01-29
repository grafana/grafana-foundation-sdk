// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Special value types supported by the `SpecialValueMap`
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum SpecialValueMatch {
    TRUE("true"),
    FALSE("false"),
    NULL("null"),
    NA_N("nan"),
    NULL_AND_NA_N("null+nan"),
    EMPTY("empty"),
    _EMPTY("");

    private final String value;

    private SpecialValueMatch(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
