// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class TextVariableBuilder implements com.grafana.foundation.cog.Builder<TextVariableKind> {
    protected final TextVariableKind internal;
    
    public TextVariableBuilder(String name) {
        this.internal = new TextVariableKind();
        this.internal.kind = "TextVariable";
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TextVariableSpec();
		}
        this.internal.spec.name = name;
    }
    public TextVariableBuilder spec(TextVariableSpec spec) {
        this.internal.spec = spec;
        return this;
    }
    
    public TextVariableBuilder name(String name) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TextVariableSpec();
		}
        this.internal.spec.name = name;
        return this;
    }
    
    public TextVariableBuilder current(VariableOption current) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TextVariableSpec();
		}
        this.internal.spec.current = current;
        return this;
    }
    
    public TextVariableBuilder query(String query) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TextVariableSpec();
		}
        this.internal.spec.query = query;
        return this;
    }
    
    public TextVariableBuilder label(String label) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TextVariableSpec();
		}
        this.internal.spec.label = label;
        return this;
    }
    
    public TextVariableBuilder hide(VariableHide hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TextVariableSpec();
		}
        this.internal.spec.hide = hide;
        return this;
    }
    
    public TextVariableBuilder skipUrlSync(Boolean skipUrlSync) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TextVariableSpec();
		}
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }
    
    public TextVariableBuilder description(String description) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TextVariableSpec();
		}
        this.internal.spec.description = description;
        return this;
    }
    public TextVariableKind build() {
        return this.internal;
    }
}
