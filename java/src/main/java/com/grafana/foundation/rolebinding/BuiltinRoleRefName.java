// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.rolebinding;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum BuiltinRoleRefName {
    VIEWER("viewer"),
    EDITOR("editor"),
    ADMIN("admin"),
    _EMPTY("");

    private final String value;

    private BuiltinRoleRefName(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
