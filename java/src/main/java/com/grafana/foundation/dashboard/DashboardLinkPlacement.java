// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Dashboard Link placement. Defines where the link should be displayed.
// - "inControlsMenu" renders the link in bottom part of the dashboard controls dropdown menu
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum DashboardLinkPlacement {
    IN_CONTROLS_MENU("inControlsMenu"),
    _EMPTY("");

    private final String value;

    private DashboardLinkPlacement(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
