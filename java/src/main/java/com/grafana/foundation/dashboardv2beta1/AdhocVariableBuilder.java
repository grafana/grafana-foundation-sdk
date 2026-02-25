// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.LinkedList;

public class AdhocVariableBuilder implements com.grafana.foundation.cog.Builder<AdhocVariableKind> {
    protected final AdhocVariableKind internal;
    
    public AdhocVariableBuilder(String name) {
        this.internal = new AdhocVariableKind();
        this.internal.kind = "AdhocVariable";
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AdhocVariableSpec();
		}
        this.internal.spec.name = name;
    }
    public AdhocVariableBuilder group(String group) {
        this.internal.group = group;
        return this;
    }
    
    public AdhocVariableBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1AdhocVariableKindDatasource> datasource) {
    Dashboardv2beta1AdhocVariableKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public AdhocVariableBuilder name(String name) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AdhocVariableSpec();
		}
        this.internal.spec.name = name;
        return this;
    }
    
    public AdhocVariableBuilder baseFilters(List<com.grafana.foundation.cog.Builder<AdHocFilterWithLabels>> baseFilters) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AdhocVariableSpec();
		}
        List<AdHocFilterWithLabels> baseFiltersResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<AdHocFilterWithLabels> r1 : baseFilters) {
                AdHocFilterWithLabels baseFiltersDepth1 = r1.build();
                baseFiltersResources.add(baseFiltersDepth1); 
        }
        this.internal.spec.baseFilters = baseFiltersResources;
        return this;
    }
    
    public AdhocVariableBuilder filters(List<com.grafana.foundation.cog.Builder<AdHocFilterWithLabels>> filters) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AdhocVariableSpec();
		}
        List<AdHocFilterWithLabels> filtersResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<AdHocFilterWithLabels> r1 : filters) {
                AdHocFilterWithLabels filtersDepth1 = r1.build();
                filtersResources.add(filtersDepth1); 
        }
        this.internal.spec.filters = filtersResources;
        return this;
    }
    
    public AdhocVariableBuilder defaultKeys(List<com.grafana.foundation.cog.Builder<MetricFindValue>> defaultKeys) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AdhocVariableSpec();
		}
        List<MetricFindValue> defaultKeysResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<MetricFindValue> r1 : defaultKeys) {
                MetricFindValue defaultKeysDepth1 = r1.build();
                defaultKeysResources.add(defaultKeysDepth1); 
        }
        this.internal.spec.defaultKeys = defaultKeysResources;
        return this;
    }
    
    public AdhocVariableBuilder label(String label) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AdhocVariableSpec();
		}
        this.internal.spec.label = label;
        return this;
    }
    
    public AdhocVariableBuilder hide(VariableHide hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AdhocVariableSpec();
		}
        this.internal.spec.hide = hide;
        return this;
    }
    
    public AdhocVariableBuilder skipUrlSync(Boolean skipUrlSync) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AdhocVariableSpec();
		}
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }
    
    public AdhocVariableBuilder description(String description) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AdhocVariableSpec();
		}
        this.internal.spec.description = description;
        return this;
    }
    
    public AdhocVariableBuilder allowCustomValue(Boolean allowCustomValue) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AdhocVariableSpec();
		}
        this.internal.spec.allowCustomValue = allowCustomValue;
        return this;
    }
    public AdhocVariableKind build() {
        return this.internal;
    }
}
