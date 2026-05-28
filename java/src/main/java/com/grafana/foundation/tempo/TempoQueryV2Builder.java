// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.tempo;

import com.grafana.foundation.dashboardv2.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2.Dashboardv2DataQueryKindDatasource;
import java.util.List;
import java.util.LinkedList;

public class TempoQueryV2Builder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public TempoQueryV2Builder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "tempo";
    }
    public TempoQueryV2Builder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public TempoQueryV2Builder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public TempoQueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource) {
    Dashboardv2DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public TempoQueryV2Builder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public TempoQueryV2Builder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public TempoQueryV2Builder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public TempoQueryV2Builder query(String query) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).query = query;
        return this;
    }
    
    public TempoQueryV2Builder search(String search) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).search = search;
        return this;
    }
    
    public TempoQueryV2Builder serviceName(String serviceName) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).serviceName = serviceName;
        return this;
    }
    
    public TempoQueryV2Builder spanName(String spanName) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).spanName = spanName;
        return this;
    }
    
    public TempoQueryV2Builder minDuration(String minDuration) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).minDuration = minDuration;
        return this;
    }
    
    public TempoQueryV2Builder maxDuration(String maxDuration) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).maxDuration = maxDuration;
        return this;
    }
    
    public TempoQueryV2Builder serviceMapQuery(StringOrArrayOfString serviceMapQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).serviceMapQuery = serviceMapQuery;
        return this;
    }
    
    public TempoQueryV2Builder serviceMapIncludeNamespace(Boolean serviceMapIncludeNamespace) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).serviceMapIncludeNamespace = serviceMapIncludeNamespace;
        return this;
    }
    
    public TempoQueryV2Builder limit(Long limit) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).limit = limit;
        return this;
    }
    
    public TempoQueryV2Builder spss(Long spss) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).spss = spss;
        return this;
    }
    
    public TempoQueryV2Builder filters(List<com.grafana.foundation.cog.Builder<TraceqlFilter>> filters) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.DataqueryBuilder().build();
		}
        List<TraceqlFilter> filtersResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<TraceqlFilter> r1 : filters) {
                TraceqlFilter filtersDepth1 = r1.build();
                filtersResources.add(filtersDepth1); 
        }
        ((Dataquery) this.internal.spec).filters = filtersResources;
        return this;
    }
    
    public TempoQueryV2Builder groupBy(List<com.grafana.foundation.cog.Builder<TraceqlFilter>> groupBy) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.DataqueryBuilder().build();
		}
        List<TraceqlFilter> groupByResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<TraceqlFilter> r1 : groupBy) {
                TraceqlFilter groupByDepth1 = r1.build();
                groupByResources.add(groupByDepth1); 
        }
        ((Dataquery) this.internal.spec).groupBy = groupByResources;
        return this;
    }
    
    public TempoQueryV2Builder tableType(SearchTableType tableType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).tableType = tableType;
        return this;
    }
    
    public TempoQueryV2Builder step(String step) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).step = step;
        return this;
    }
    
    public TempoQueryV2Builder exemplars(Long exemplars) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).exemplars = exemplars;
        return this;
    }
    
    public TempoQueryV2Builder metricsQueryType(MetricsQueryType metricsQueryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.tempo.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).metricsQueryType = metricsQueryType;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
