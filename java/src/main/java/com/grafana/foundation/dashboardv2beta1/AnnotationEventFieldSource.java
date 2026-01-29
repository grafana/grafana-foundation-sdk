// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Annotation event field source. Defines how to obtain the value for an annotation event field.
// - "field": Find the value with a matching key (default)
// - "text": Write a constant string into the value
// - "skip": Do not include the field
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum AnnotationEventFieldSource {
    FIELD("field"),
    TEXT("text"),
    SKIP("skip"),
    _EMPTY("");

    private final String value;

    private AnnotationEventFieldSource(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
