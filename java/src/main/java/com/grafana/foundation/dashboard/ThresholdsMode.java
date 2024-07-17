// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Thresholds can either be `absolute` (specific number) or `percentage` (relative to min or max, it will be values between 0 and 1).
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ThresholdsMode {
    ABSOLUTE("absolute"),
    PERCENTAGE("percentage"),
    _EMPTY("");

    private final String value;

    private ThresholdsMode(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
