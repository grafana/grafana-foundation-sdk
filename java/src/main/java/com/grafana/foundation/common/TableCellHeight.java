// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Height of a table cell
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum TableCellHeight {
    SM("sm"),
    MD("md"),
    LG("lg"),
    AUTO("auto"),
    _EMPTY("");

    private final String value;

    private TableCellHeight(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
