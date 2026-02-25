// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class SwitchVariableBuilder implements com.grafana.foundation.cog.Builder<SwitchVariableKind> {
    protected final SwitchVariableKind internal;
    
    public SwitchVariableBuilder(String name) {
        this.internal = new SwitchVariableKind();
        this.internal.kind = "SwitchVariable";
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.SwitchVariableSpec();
		}
        this.internal.spec.name = name;
    }
    public SwitchVariableBuilder name(String name) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.SwitchVariableSpec();
		}
        this.internal.spec.name = name;
        return this;
    }
    
    public SwitchVariableBuilder current(String current) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.SwitchVariableSpec();
		}
        this.internal.spec.current = current;
        return this;
    }
    
    public SwitchVariableBuilder enabledValue(String enabledValue) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.SwitchVariableSpec();
		}
        this.internal.spec.enabledValue = enabledValue;
        return this;
    }
    
    public SwitchVariableBuilder disabledValue(String disabledValue) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.SwitchVariableSpec();
		}
        this.internal.spec.disabledValue = disabledValue;
        return this;
    }
    
    public SwitchVariableBuilder label(String label) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.SwitchVariableSpec();
		}
        this.internal.spec.label = label;
        return this;
    }
    
    public SwitchVariableBuilder hide(VariableHide hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.SwitchVariableSpec();
		}
        this.internal.spec.hide = hide;
        return this;
    }
    
    public SwitchVariableBuilder skipUrlSync(Boolean skipUrlSync) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.SwitchVariableSpec();
		}
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }
    
    public SwitchVariableBuilder description(String description) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.SwitchVariableSpec();
		}
        this.internal.spec.description = description;
        return this;
    }
    public SwitchVariableKind build() {
        return this.internal;
    }
}
