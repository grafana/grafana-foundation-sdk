// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.playlistv0alpha1;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ItemType {
    DASHBOARD_BY_TAG("dashboard_by_tag"),
    DASHBOARD_BY_UID("dashboard_by_uid"),
    DASHBOARD_BY_ID("dashboard_by_id"),
    _EMPTY("");

    private final String value;

    private ItemType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
