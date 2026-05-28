// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2;


public class TextVariableBuilder implements com.grafana.foundation.cog.Builder<TextVariableKind> {
    protected final TextVariableKind internal;
    
    public TextVariableBuilder(String name) {
        this.internal = new TextVariableKind();
        this.internal.kind = "TextVariable";
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.TextVariableSpec();
		}
        this.internal.spec.name = name;
    }
    public TextVariableBuilder name(String name) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.TextVariableSpec();
		}
        this.internal.spec.name = name;
        return this;
    }
    
    public TextVariableBuilder current(VariableOption current) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.TextVariableSpec();
		}
        this.internal.spec.current = current;
        return this;
    }
    
    public TextVariableBuilder query(String query) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.TextVariableSpec();
		}
        this.internal.spec.query = query;
        return this;
    }
    
    public TextVariableBuilder label(String label) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.TextVariableSpec();
		}
        this.internal.spec.label = label;
        return this;
    }
    
    public TextVariableBuilder hide(VariableHide hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.TextVariableSpec();
		}
        this.internal.spec.hide = hide;
        return this;
    }
    
    public TextVariableBuilder skipUrlSync(Boolean skipUrlSync) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.TextVariableSpec();
		}
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }
    
    public TextVariableBuilder description(String description) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.TextVariableSpec();
		}
        this.internal.spec.description = description;
        return this;
    }
    
    public TextVariableBuilder origin(com.grafana.foundation.cog.Builder<ControlSourceRef> origin) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.TextVariableSpec();
		}
    ControlSourceRef originResource = origin.build();
        this.internal.spec.origin = originResource;
        return this;
    }
    public TextVariableKind build() {
        return this.internal;
    }
}
