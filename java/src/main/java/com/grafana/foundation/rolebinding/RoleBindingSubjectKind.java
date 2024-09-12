// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.rolebinding;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum RoleBindingSubjectKind {
    TEAM("Team"),
    USER("User"),
    _EMPTY("");

    private final String value;

    private RoleBindingSubjectKind(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
