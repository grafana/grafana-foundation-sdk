// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum BuilderQueryEditorOrderByOptions {
    ASC("asc"),
    DESC("desc"),
    _EMPTY("");

    private final String value;

    private BuilderQueryEditorOrderByOptions(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
