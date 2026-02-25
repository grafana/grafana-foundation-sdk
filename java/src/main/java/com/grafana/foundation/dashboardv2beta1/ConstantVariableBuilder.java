// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class ConstantVariableBuilder implements com.grafana.foundation.cog.Builder<ConstantVariableKind> {
    protected final ConstantVariableKind internal;
    
    public ConstantVariableBuilder(String name) {
        this.internal = new ConstantVariableKind();
        this.internal.kind = "ConstantVariable";
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.ConstantVariableSpec();
		}
        this.internal.spec.name = name;
    }
    public ConstantVariableBuilder name(String name) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.ConstantVariableSpec();
		}
        this.internal.spec.name = name;
        return this;
    }
    
    public ConstantVariableBuilder query(String query) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.ConstantVariableSpec();
		}
        this.internal.spec.query = query;
        return this;
    }
    
    public ConstantVariableBuilder current(VariableOption current) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.ConstantVariableSpec();
		}
        this.internal.spec.current = current;
        return this;
    }
    
    public ConstantVariableBuilder label(String label) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.ConstantVariableSpec();
		}
        this.internal.spec.label = label;
        return this;
    }
    
    public ConstantVariableBuilder hide(VariableHide hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.ConstantVariableSpec();
		}
        this.internal.spec.hide = hide;
        return this;
    }
    
    public ConstantVariableBuilder skipUrlSync(Boolean skipUrlSync) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.ConstantVariableSpec();
		}
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }
    
    public ConstantVariableBuilder description(String description) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.ConstantVariableSpec();
		}
        this.internal.spec.description = description;
        return this;
    }
    public ConstantVariableKind build() {
        return this.internal;
    }
}
