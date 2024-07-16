// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum MetricFindQueryTypes {
    PROJECTS("projects"),
    SERVICES("services"),
    DEFAULT_PROJECT("defaultProject"),
    METRIC_TYPES("metricTypes"),
    LABEL_KEYS("labelKeys"),
    LABEL_VALUES("labelValues"),
    RESOURCE_TYPES("resourceTypes"),
    AGGREGATIONS("aggregations"),
    ALIGNERS("aligners"),
    ALIGNMENT_PERIODS("alignmentPeriods"),
    SELECTORS("selectors"),
    SLO_SERVICES("sloServices"),
    SLO("slo"),
    _EMPTY("");

    private final String value;

    private MetricFindQueryTypes(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
