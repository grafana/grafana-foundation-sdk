// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum AlignmentTypes {
    ALIGN_DELTA("ALIGN_DELTA"),
    ALIGN_RATE("ALIGN_RATE"),
    ALIGN_INTERPOLATE("ALIGN_INTERPOLATE"),
    ALIGN_NEXT_OLDER("ALIGN_NEXT_OLDER"),
    ALIGN_MIN("ALIGN_MIN"),
    ALIGN_MAX("ALIGN_MAX"),
    ALIGN_MEAN("ALIGN_MEAN"),
    ALIGN_COUNT("ALIGN_COUNT"),
    ALIGN_SUM("ALIGN_SUM"),
    ALIGN_STDDEV("ALIGN_STDDEV"),
    ALIGN_COUNT_TRUE("ALIGN_COUNT_TRUE"),
    ALIGN_COUNT_FALSE("ALIGN_COUNT_FALSE"),
    ALIGN_FRACTION_TRUE("ALIGN_FRACTION_TRUE"),
    ALIGN_PERCENTILE_99("ALIGN_PERCENTILE_99"),
    ALIGN_PERCENTILE_95("ALIGN_PERCENTILE_95"),
    ALIGN_PERCENTILE_50("ALIGN_PERCENTILE_50"),
    ALIGN_PERCENTILE_05("ALIGN_PERCENTILE_05"),
    ALIGN_PERCENT_CHANGE("ALIGN_PERCENT_CHANGE"),
    ALIGN_NONE("ALIGN_NONE"),
    _EMPTY("");

    private final String value;

    private AlignmentTypes(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
