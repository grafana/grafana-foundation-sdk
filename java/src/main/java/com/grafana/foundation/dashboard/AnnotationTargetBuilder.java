// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.List;

public class AnnotationTargetBuilder implements com.grafana.foundation.cog.Builder<AnnotationTarget> {
    protected final AnnotationTarget internal;
    
    public AnnotationTargetBuilder() {
        this.internal = new AnnotationTarget();
    }
    public AnnotationTargetBuilder limit(Long limit) {
        this.internal.limit = limit;
        return this;
    }
    
    public AnnotationTargetBuilder matchAny(Boolean matchAny) {
        this.internal.matchAny = matchAny;
        return this;
    }
    
    public AnnotationTargetBuilder tags(List<String> tags) {
        this.internal.tags = tags;
        return this;
    }
    
    public AnnotationTargetBuilder type(String type) {
        this.internal.type = type;
        return this;
    }
    public AnnotationTarget build() {
        return this.internal;
    }
}
