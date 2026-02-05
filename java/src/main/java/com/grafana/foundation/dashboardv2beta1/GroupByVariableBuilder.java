// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;

public class GroupByVariableBuilder implements com.grafana.foundation.cog.Builder<GroupByVariableKind> {
    protected final GroupByVariableKind internal;
    
    public GroupByVariableBuilder(String name) {
        this.internal = new GroupByVariableKind();
        this.internal.kind = "GroupByVariable";
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.GroupByVariableSpec();
		}
        this.internal.spec.name = name;
    }
    public GroupByVariableBuilder group(String group) {
        this.internal.group = group;
        return this;
    }
    
    public GroupByVariableBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1GroupByVariableKindDatasource> datasource) {
    Dashboardv2beta1GroupByVariableKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public GroupByVariableBuilder spec(GroupByVariableSpec spec) {
        this.internal.spec = spec;
        return this;
    }
    
    public GroupByVariableBuilder name(String name) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.GroupByVariableSpec();
		}
        this.internal.spec.name = name;
        return this;
    }
    
    public GroupByVariableBuilder defaultValue(VariableOption defaultValue) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.GroupByVariableSpec();
		}
        this.internal.spec.defaultValue = defaultValue;
        return this;
    }
    
    public GroupByVariableBuilder current(VariableOption current) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.GroupByVariableSpec();
		}
        this.internal.spec.current = current;
        return this;
    }
    
    public GroupByVariableBuilder options(List<VariableOption> options) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.GroupByVariableSpec();
		}
        this.internal.spec.options = options;
        return this;
    }
    
    public GroupByVariableBuilder multi(Boolean multi) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.GroupByVariableSpec();
		}
        this.internal.spec.multi = multi;
        return this;
    }
    
    public GroupByVariableBuilder label(String label) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.GroupByVariableSpec();
		}
        this.internal.spec.label = label;
        return this;
    }
    
    public GroupByVariableBuilder hide(VariableHide hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.GroupByVariableSpec();
		}
        this.internal.spec.hide = hide;
        return this;
    }
    
    public GroupByVariableBuilder skipUrlSync(Boolean skipUrlSync) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.GroupByVariableSpec();
		}
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }
    
    public GroupByVariableBuilder description(String description) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.GroupByVariableSpec();
		}
        this.internal.spec.description = description;
        return this;
    }
    public GroupByVariableKind build() {
        return this.internal;
    }
}
