// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;

public class CustomVariableBuilder implements com.grafana.foundation.cog.Builder<CustomVariableKind> {
    protected final CustomVariableKind internal;
    
    public CustomVariableBuilder(String name) {
        this.internal = new CustomVariableKind();
        this.internal.kind = "CustomVariable";
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.CustomVariableSpec();
		}
        this.internal.spec.name = name;
    }
    public CustomVariableBuilder name(String name) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.CustomVariableSpec();
		}
        this.internal.spec.name = name;
        return this;
    }
    
    public CustomVariableBuilder query(String query) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.CustomVariableSpec();
		}
        this.internal.spec.query = query;
        return this;
    }
    
    public CustomVariableBuilder current(VariableOption current) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.CustomVariableSpec();
		}
        this.internal.spec.current = current;
        return this;
    }
    
    public CustomVariableBuilder options(List<VariableOption> options) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.CustomVariableSpec();
		}
        this.internal.spec.options = options;
        return this;
    }
    
    public CustomVariableBuilder multi(Boolean multi) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.CustomVariableSpec();
		}
        this.internal.spec.multi = multi;
        return this;
    }
    
    public CustomVariableBuilder includeAll(Boolean includeAll) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.CustomVariableSpec();
		}
        this.internal.spec.includeAll = includeAll;
        return this;
    }
    
    public CustomVariableBuilder allValue(String allValue) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.CustomVariableSpec();
		}
        this.internal.spec.allValue = allValue;
        return this;
    }
    
    public CustomVariableBuilder label(String label) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.CustomVariableSpec();
		}
        this.internal.spec.label = label;
        return this;
    }
    
    public CustomVariableBuilder hide(VariableHide hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.CustomVariableSpec();
		}
        this.internal.spec.hide = hide;
        return this;
    }
    
    public CustomVariableBuilder skipUrlSync(Boolean skipUrlSync) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.CustomVariableSpec();
		}
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }
    
    public CustomVariableBuilder description(String description) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.CustomVariableSpec();
		}
        this.internal.spec.description = description;
        return this;
    }
    
    public CustomVariableBuilder allowCustomValue(Boolean allowCustomValue) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.CustomVariableSpec();
		}
        this.internal.spec.allowCustomValue = allowCustomValue;
        return this;
    }
    public CustomVariableKind build() {
        return this.internal;
    }
}
