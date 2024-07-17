// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.accesspolicy;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum RoleRefKind {
    ROLE("Role"),
    BUILTIN_ROLE("BuiltinRole"),
    TEAM("Team"),
    USER("User"),
    _EMPTY("");

    private final String value;

    private RoleRefKind(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
