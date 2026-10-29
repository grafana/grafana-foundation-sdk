// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum NotificationSettingsType {
    SIMPLIFIED_ROUTING("SimplifiedRouting"),
    NAMED_ROUTING_TREE("NamedRoutingTree"),
    _EMPTY("");

    private final String value;

    private NotificationSettingsType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
