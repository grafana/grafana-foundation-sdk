// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum TypeResampleUpsampler {
    PAD("pad"),
    BACKFILLING("backfilling"),
    FILLNA("fillna"),
    _EMPTY("");

    private final String value;

    private TypeResampleUpsampler(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
