// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum DataqueryScenarioId {
    ANNOTATIONS("annotations"),
    ARROW("arrow"),
    CSV_CONTENT("csv_content"),
    CSV_FILE("csv_file"),
    CSV_METRIC_VALUES("csv_metric_values"),
    DATAPOINTS_OUTSIDE_RANGE("datapoints_outside_range"),
    ERROR_WITH_SOURCE("error_with_source"),
    EXPONENTIAL_HEATMAP_BUCKET_DATA("exponential_heatmap_bucket_data"),
    FLAME_GRAPH("flame_graph"),
    GRAFANA_API("grafana_api"),
    LINEAR_HEATMAP_BUCKET_DATA("linear_heatmap_bucket_data"),
    LIVE("live"),
    LOGS("logs"),
    MANUAL_ENTRY("manual_entry"),
    NO_DATA_POINTS("no_data_points"),
    NODE_GRAPH("node_graph"),
    PREDICTABLE_CSV_WAVE("predictable_csv_wave"),
    PREDICTABLE_PULSE("predictable_pulse"),
    RANDOM_WALK("random_walk"),
    RANDOM_WALK_TABLE("random_walk_table"),
    RANDOM_WALK_WITH_ERROR("random_walk_with_error"),
    RAW_FRAME("raw_frame"),
    SERVER_ERROR500("server_error_500"),
    SIMULATION("simulation"),
    SLOW_QUERY("slow_query"),
    STREAMING_CLIENT("streaming_client"),
    TABLE_STATIC("table_static"),
    TRACE("trace"),
    USA("usa"),
    VARIABLES_QUERY("variables-query"),
    _EMPTY("");

    private final String value;

    private DataqueryScenarioId(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
