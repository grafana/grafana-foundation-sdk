// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class AnnotationEventFieldMappingBuilder implements com.grafana.foundation.cog.Builder<AnnotationEventFieldMapping> {
    protected final AnnotationEventFieldMapping internal;
    
    public AnnotationEventFieldMappingBuilder() {
        this.internal = new AnnotationEventFieldMapping();
    }
    public AnnotationEventFieldMappingBuilder source(String source) {
        this.internal.source = source;
        return this;
    }
    
    public AnnotationEventFieldMappingBuilder value(String value) {
        this.internal.value = value;
        return this;
    }
    
    public AnnotationEventFieldMappingBuilder regex(String regex) {
        this.internal.regex = regex;
        return this;
    }
    public AnnotationEventFieldMapping build() {
        return this.internal;
    }
}
