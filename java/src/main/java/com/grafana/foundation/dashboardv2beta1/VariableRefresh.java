// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Options to config when to refresh a variable
// `never`: Never refresh the variable
// `onDashboardLoad`: Queries the data source every time the dashboard loads.
// `onTimeRangeChanged`: Queries the data source when the dashboard time range changes.
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum VariableRefresh {
    NEVER("never"),
    ON_DASHBOARD_LOAD("onDashboardLoad"),
    ON_TIME_RANGE_CHANGED("onTimeRangeChanged"),
    _EMPTY("");

    private final String value;

    private VariableRefresh(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
