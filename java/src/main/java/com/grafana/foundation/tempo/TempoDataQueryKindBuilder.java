// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.tempo;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;
import java.util.List;
import java.util.LinkedList;

public class TempoDataQueryKindBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public TempoDataQueryKindBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "tempo";
    }
    public TempoDataQueryKindBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public TempoDataQueryKindBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public TempoDataQueryKindBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.TempoQueryBuilder().build();
		}
        ((TempoQuery) this.internal.spec).refId = refId;
        return this;
    }
    
    public TempoDataQueryKindBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.TempoQueryBuilder().build();
		}
        ((TempoQuery) this.internal.spec).hide = hide;
        return this;
    }
    
    public TempoDataQueryKindBuilder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.TempoQueryBuilder().build();
		}
        ((TempoQuery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public TempoDataQueryKindBuilder query(String query) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.TempoQueryBuilder().build();
		}
        ((TempoQuery) this.internal.spec).query = query;
        return this;
    }
    
    public TempoDataQueryKindBuilder search(String search) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.TempoQueryBuilder().build();
		}
        ((TempoQuery) this.internal.spec).search = search;
        return this;
    }
    
    public TempoDataQueryKindBuilder serviceName(String serviceName) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.TempoQueryBuilder().build();
		}
        ((TempoQuery) this.internal.spec).serviceName = serviceName;
        return this;
    }
    
    public TempoDataQueryKindBuilder spanName(String spanName) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.TempoQueryBuilder().build();
		}
        ((TempoQuery) this.internal.spec).spanName = spanName;
        return this;
    }
    
    public TempoDataQueryKindBuilder minDuration(String minDuration) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.TempoQueryBuilder().build();
		}
        ((TempoQuery) this.internal.spec).minDuration = minDuration;
        return this;
    }
    
    public TempoDataQueryKindBuilder maxDuration(String maxDuration) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.TempoQueryBuilder().build();
		}
        ((TempoQuery) this.internal.spec).maxDuration = maxDuration;
        return this;
    }
    
    public TempoDataQueryKindBuilder serviceMapQuery(StringOrArrayOfString serviceMapQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.TempoQueryBuilder().build();
		}
        ((TempoQuery) this.internal.spec).serviceMapQuery = serviceMapQuery;
        return this;
    }
    
    public TempoDataQueryKindBuilder serviceMapIncludeNamespace(Boolean serviceMapIncludeNamespace) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.TempoQueryBuilder().build();
		}
        ((TempoQuery) this.internal.spec).serviceMapIncludeNamespace = serviceMapIncludeNamespace;
        return this;
    }
    
    public TempoDataQueryKindBuilder limit(Long limit) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.TempoQueryBuilder().build();
		}
        ((TempoQuery) this.internal.spec).limit = limit;
        return this;
    }
    
    public TempoDataQueryKindBuilder spss(Long spss) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.TempoQueryBuilder().build();
		}
        ((TempoQuery) this.internal.spec).spss = spss;
        return this;
    }
    
    public TempoDataQueryKindBuilder filters(List<com.grafana.foundation.cog.Builder<TraceqlFilter>> filters) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.TempoQueryBuilder().build();
		}
        List<TraceqlFilter> filtersResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<TraceqlFilter> r1 : filters) {
                TraceqlFilter filtersDepth1 = r1.build();
                filtersResources.add(filtersDepth1); 
        }
        ((TempoQuery) this.internal.spec).filters = filtersResources;
        return this;
    }
    
    public TempoDataQueryKindBuilder groupBy(List<com.grafana.foundation.cog.Builder<TraceqlFilter>> groupBy) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.TempoQueryBuilder().build();
		}
        List<TraceqlFilter> groupByResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<TraceqlFilter> r1 : groupBy) {
                TraceqlFilter groupByDepth1 = r1.build();
                groupByResources.add(groupByDepth1); 
        }
        ((TempoQuery) this.internal.spec).groupBy = groupByResources;
        return this;
    }
    
    public TempoDataQueryKindBuilder tableType(SearchTableType tableType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.TempoQueryBuilder().build();
		}
        ((TempoQuery) this.internal.spec).tableType = tableType;
        return this;
    }
    
    public TempoDataQueryKindBuilder step(String step) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.TempoQueryBuilder().build();
		}
        ((TempoQuery) this.internal.spec).step = step;
        return this;
    }
    
    public TempoDataQueryKindBuilder exemplars(Long exemplars) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.TempoQueryBuilder().build();
		}
        ((TempoQuery) this.internal.spec).exemplars = exemplars;
        return this;
    }
    
    public TempoDataQueryKindBuilder metricsQueryType(MetricsQueryType metricsQueryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.TempoQueryBuilder().build();
		}
        ((TempoQuery) this.internal.spec).metricsQueryType = metricsQueryType;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
