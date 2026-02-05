// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;

public class AnnotationPanelFilterBuilder implements com.grafana.foundation.cog.Builder<AnnotationPanelFilter> {
    protected final AnnotationPanelFilter internal;
    
    public AnnotationPanelFilterBuilder() {
        this.internal = new AnnotationPanelFilter();
    }
    public AnnotationPanelFilterBuilder exclude(Boolean exclude) {
        this.internal.exclude = exclude;
        return this;
    }
    
    public AnnotationPanelFilterBuilder ids(List<Integer> ids) {
        this.internal.ids = ids;
        return this;
    }
    public AnnotationPanelFilter build() {
        return this.internal;
    }
}
