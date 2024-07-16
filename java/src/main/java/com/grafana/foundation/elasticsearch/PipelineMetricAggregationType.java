// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum PipelineMetricAggregationType {
    MOVING_AVG("moving_avg"),
    MOVING_FN("moving_fn"),
    DERIVATIVE("derivative"),
    SERIAL_DIFF("serial_diff"),
    CUMULATIVE_SUM("cumulative_sum"),
    BUCKET_SCRIPT("bucket_script"),
    _EMPTY("");

    private final String value;

    private PipelineMetricAggregationType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
