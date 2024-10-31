// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.librarypanel;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum LibraryPanelRepeatDirection {
    H("h"),
    V("v"),
    _EMPTY("");

    private final String value;

    private LibraryPanelRepeatDirection(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
