// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;

public class DatasourceVariableBuilder implements com.grafana.foundation.cog.Builder<DatasourceVariableKind> {
    protected final DatasourceVariableKind internal;
    
    public DatasourceVariableBuilder(String name) {
        this.internal = new DatasourceVariableKind();
        this.internal.kind = "DatasourceVariable";
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.DatasourceVariableSpec();
		}
        this.internal.spec.name = name;
    }
    public DatasourceVariableBuilder name(String name) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.DatasourceVariableSpec();
		}
        this.internal.spec.name = name;
        return this;
    }
    
    public DatasourceVariableBuilder pluginId(String pluginId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.DatasourceVariableSpec();
		}
        this.internal.spec.pluginId = pluginId;
        return this;
    }
    
    public DatasourceVariableBuilder refresh(VariableRefresh refresh) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.DatasourceVariableSpec();
		}
        this.internal.spec.refresh = refresh;
        return this;
    }
    
    public DatasourceVariableBuilder regex(String regex) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.DatasourceVariableSpec();
		}
        this.internal.spec.regex = regex;
        return this;
    }
    
    public DatasourceVariableBuilder current(VariableOption current) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.DatasourceVariableSpec();
		}
        this.internal.spec.current = current;
        return this;
    }
    
    public DatasourceVariableBuilder options(List<VariableOption> options) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.DatasourceVariableSpec();
		}
        this.internal.spec.options = options;
        return this;
    }
    
    public DatasourceVariableBuilder multi(Boolean multi) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.DatasourceVariableSpec();
		}
        this.internal.spec.multi = multi;
        return this;
    }
    
    public DatasourceVariableBuilder includeAll(Boolean includeAll) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.DatasourceVariableSpec();
		}
        this.internal.spec.includeAll = includeAll;
        return this;
    }
    
    public DatasourceVariableBuilder allValue(String allValue) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.DatasourceVariableSpec();
		}
        this.internal.spec.allValue = allValue;
        return this;
    }
    
    public DatasourceVariableBuilder label(String label) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.DatasourceVariableSpec();
		}
        this.internal.spec.label = label;
        return this;
    }
    
    public DatasourceVariableBuilder hide(VariableHide hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.DatasourceVariableSpec();
		}
        this.internal.spec.hide = hide;
        return this;
    }
    
    public DatasourceVariableBuilder skipUrlSync(Boolean skipUrlSync) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.DatasourceVariableSpec();
		}
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }
    
    public DatasourceVariableBuilder description(String description) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.DatasourceVariableSpec();
		}
        this.internal.spec.description = description;
        return this;
    }
    
    public DatasourceVariableBuilder allowCustomValue(Boolean allowCustomValue) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.DatasourceVariableSpec();
		}
        this.internal.spec.allowCustomValue = allowCustomValue;
        return this;
    }
    public DatasourceVariableKind build() {
        return this.internal;
    }
}
