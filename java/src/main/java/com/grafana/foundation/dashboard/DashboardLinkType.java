// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Dashboard Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum DashboardLinkType {
    LINK("link"),
    DASHBOARDS("dashboards"),
    _EMPTY("");

    private final String value;

    private DashboardLinkType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
