// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// 0 for no shared crosshair or tooltip (default).
// 1 for shared crosshair.
// 2 for shared crosshair AND shared tooltip.
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum DashboardCursorSync {
    OFF(0),
    CROSSHAIR(1),
    TOOLTIP(2);

    private final Integer value;

    private DashboardCursorSync(Integer value) {
        this.value = value;
    }

    @JsonValue
    public Integer Value() {
        return value;
    }
}
