// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;

public class QueryVariableBuilder implements com.grafana.foundation.cog.Builder<QueryVariableKind> {
    protected final QueryVariableKind internal;
    
    public QueryVariableBuilder(String name) {
        this.internal = new QueryVariableKind();
        this.internal.kind = "QueryVariable";
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
        this.internal.spec.name = name;
    }
    public QueryVariableBuilder name(String name) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
        this.internal.spec.name = name;
        return this;
    }
    
    public QueryVariableBuilder current(VariableOption current) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
        this.internal.spec.current = current;
        return this;
    }
    
    public QueryVariableBuilder label(String label) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
        this.internal.spec.label = label;
        return this;
    }
    
    public QueryVariableBuilder hide(VariableHide hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
        this.internal.spec.hide = hide;
        return this;
    }
    
    public QueryVariableBuilder refresh(VariableRefresh refresh) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
        this.internal.spec.refresh = refresh;
        return this;
    }
    
    public QueryVariableBuilder skipUrlSync(Boolean skipUrlSync) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }
    
    public QueryVariableBuilder description(String description) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
        this.internal.spec.description = description;
        return this;
    }
    
    public QueryVariableBuilder query(com.grafana.foundation.cog.Builder<DataQueryKind> query) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
    DataQueryKind queryResource = query.build();
        this.internal.spec.query = queryResource;
        return this;
    }
    
    public QueryVariableBuilder regex(String regex) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
        this.internal.spec.regex = regex;
        return this;
    }
    
    public QueryVariableBuilder sort(VariableSort sort) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
        this.internal.spec.sort = sort;
        return this;
    }
    
    public QueryVariableBuilder definition(String definition) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
        this.internal.spec.definition = definition;
        return this;
    }
    
    public QueryVariableBuilder options(List<VariableOption> options) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
        this.internal.spec.options = options;
        return this;
    }
    
    public QueryVariableBuilder multi(Boolean multi) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
        this.internal.spec.multi = multi;
        return this;
    }
    
    public QueryVariableBuilder includeAll(Boolean includeAll) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
        this.internal.spec.includeAll = includeAll;
        return this;
    }
    
    public QueryVariableBuilder allValue(String allValue) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
        this.internal.spec.allValue = allValue;
        return this;
    }
    
    public QueryVariableBuilder placeholder(String placeholder) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
        this.internal.spec.placeholder = placeholder;
        return this;
    }
    
    public QueryVariableBuilder allowCustomValue(Boolean allowCustomValue) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
        this.internal.spec.allowCustomValue = allowCustomValue;
        return this;
    }
    
    public QueryVariableBuilder staticOptions(List<VariableOption> staticOptions) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
        this.internal.spec.staticOptions = staticOptions;
        return this;
    }
    
    public QueryVariableBuilder staticOptionsOrder(QueryVariableSpecStaticOptionsOrder staticOptionsOrder) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryVariableSpec();
		}
        this.internal.spec.staticOptionsOrder = staticOptionsOrder;
        return this;
    }
    public QueryVariableKind build() {
        return this.internal;
    }
}
