// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;

public class IntervalVariableBuilder implements com.grafana.foundation.cog.Builder<IntervalVariableKind> {
    protected final IntervalVariableKind internal;
    
    public IntervalVariableBuilder(String name) {
        this.internal = new IntervalVariableKind();
        this.internal.kind = "IntervalVariable";
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.IntervalVariableSpec();
		}
        this.internal.spec.name = name;
    }
    public IntervalVariableBuilder spec(IntervalVariableSpec spec) {
        this.internal.spec = spec;
        return this;
    }
    
    public IntervalVariableBuilder name(String name) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.IntervalVariableSpec();
		}
        this.internal.spec.name = name;
        return this;
    }
    
    public IntervalVariableBuilder query(String query) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.IntervalVariableSpec();
		}
        this.internal.spec.query = query;
        return this;
    }
    
    public IntervalVariableBuilder current(VariableOption current) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.IntervalVariableSpec();
		}
        this.internal.spec.current = current;
        return this;
    }
    
    public IntervalVariableBuilder options(List<VariableOption> options) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.IntervalVariableSpec();
		}
        this.internal.spec.options = options;
        return this;
    }
    
    public IntervalVariableBuilder auto(Boolean auto) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.IntervalVariableSpec();
		}
        this.internal.spec.auto = auto;
        return this;
    }
    
    public IntervalVariableBuilder autoMin(String autoMin) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.IntervalVariableSpec();
		}
        this.internal.spec.autoMin = autoMin;
        return this;
    }
    
    public IntervalVariableBuilder autoCount(Long autoCount) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.IntervalVariableSpec();
		}
        this.internal.spec.autoCount = autoCount;
        return this;
    }
    
    public IntervalVariableBuilder refresh(VariableRefresh refresh) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.IntervalVariableSpec();
		}
        this.internal.spec.refresh = refresh;
        return this;
    }
    
    public IntervalVariableBuilder label(String label) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.IntervalVariableSpec();
		}
        this.internal.spec.label = label;
        return this;
    }
    
    public IntervalVariableBuilder hide(VariableHide hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.IntervalVariableSpec();
		}
        this.internal.spec.hide = hide;
        return this;
    }
    
    public IntervalVariableBuilder skipUrlSync(Boolean skipUrlSync) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.IntervalVariableSpec();
		}
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }
    
    public IntervalVariableBuilder description(String description) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.IntervalVariableSpec();
		}
        this.internal.spec.description = description;
        return this;
    }
    public IntervalVariableKind build() {
        return this.internal;
    }
}
