// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.text;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum TextMode {
    HTML("html"),
    MARKDOWN("markdown"),
    CODE("code"),
    _EMPTY("");

    private final String value;

    private TextMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
