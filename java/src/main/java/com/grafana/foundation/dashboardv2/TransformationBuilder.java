// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2;


public class TransformationBuilder implements com.grafana.foundation.cog.Builder<TransformationKind> {
    protected final TransformationKind internal;
    
    public TransformationBuilder() {
        this.internal = new TransformationKind();
        this.internal.kind = "Transformation";
    }
    public TransformationBuilder group(String group) {
        this.internal.group = group;
        return this;
    }
    
    public TransformationBuilder disabled(Boolean disabled) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.TransformationSpec();
		}
        this.internal.spec.disabled = disabled;
        return this;
    }
    
    public TransformationBuilder filter(MatcherConfig filter) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.TransformationSpec();
		}
        this.internal.spec.filter = filter;
        return this;
    }
    
    public TransformationBuilder topic(DataTopic topic) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.TransformationSpec();
		}
        this.internal.spec.topic = topic;
        return this;
    }
    
    public TransformationBuilder options(Object options) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.TransformationSpec();
		}
        this.internal.spec.options = options;
        return this;
    }
    public TransformationKind build() {
        return this.internal;
    }
}
