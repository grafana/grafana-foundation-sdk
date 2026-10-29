// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// #TotalHitsRelation says how totalHits relates to the real number of matching
// rules the caller may see: "eq" when it is exact, "lte" when it is an upper
// bound because authorisation was applied after the search ranked its results.
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum TotalHitsRelation {
    EQ("eq"),
    LTE("lte"),
    _EMPTY("");

    private final String value;

    private TotalHitsRelation(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
