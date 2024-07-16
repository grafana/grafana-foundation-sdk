// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum ContactPointType {
    ALERTMANAGER("alertmanager"),
    DINGDING(" dingding"),
    DISCORD(" discord"),
    EMAIL(" email"),
    GOOGLECHAT(" googlechat"),
    KAFKA(" kafka"),
    LINE(" line"),
    OPSGENIE(" opsgenie"),
    PAGERDUTY(" pagerduty"),
    PUSHOVER(" pushover"),
    SENSUGO(" sensugo"),
    SLACK(" slack"),
    TEAMS(" teams"),
    TELEGRAM(" telegram"),
    THREEMA(" threema"),
    VICTOROPS(" victorops"),
    WEBHOOK(" webhook"),
    WECOM(" wecom"),
    _EMPTY("");

    private final String value;

    private ContactPointType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
