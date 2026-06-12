// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.prometheus;

import com.grafana.foundation.dashboardv2.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2.Dashboardv2DataQueryKindDatasource;
import java.util.List;
import java.util.LinkedList;

public class PrometheusQueryV2Builder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public PrometheusQueryV2Builder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "prometheus";
    }
    public PrometheusQueryV2Builder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public PrometheusQueryV2Builder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public PrometheusQueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource) {
    Dashboardv2DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public PrometheusQueryV2Builder adhocFilters(List<com.grafana.foundation.cog.Builder<AdhocFilters>> adhocFilters) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        List<AdhocFilters> adhocFiltersResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<AdhocFilters> r1 : adhocFilters) {
                AdhocFilters adhocFiltersDepth1 = r1.build();
                adhocFiltersResources.add(adhocFiltersDepth1); 
        }
        ((Dataquery) this.internal.spec).adhocFilters = adhocFiltersResources;
        return this;
    }
    
    public PrometheusQueryV2Builder editorMode(QueryEditorMode editorMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).editorMode = editorMode;
        return this;
    }
    
    public PrometheusQueryV2Builder exemplar(Boolean exemplar) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).exemplar = exemplar;
        return this;
    }
    
    public PrometheusQueryV2Builder expr(String expr) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).expr = expr;
        return this;
    }
    
    public PrometheusQueryV2Builder format(PromQueryFormat format) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).format = format;
        return this;
    }
    
    public PrometheusQueryV2Builder groupByKeys(List<String> groupByKeys) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).groupByKeys = groupByKeys;
        return this;
    }
    
    public PrometheusQueryV2Builder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public PrometheusQueryV2Builder instant(Boolean instant) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).instant = instant;
        return this;
    }
    
    public PrometheusQueryV2Builder intervalFactor(Long intervalFactor) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).intervalFactor = intervalFactor;
        return this;
    }
    
    public PrometheusQueryV2Builder intervalMs(Double intervalMs) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).intervalMs = intervalMs;
        return this;
    }
    
    public PrometheusQueryV2Builder legendFormat(String legendFormat) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).legendFormat = legendFormat;
        return this;
    }
    
    public PrometheusQueryV2Builder maxDataPoints(Long maxDataPoints) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).maxDataPoints = maxDataPoints;
        return this;
    }
    
    public PrometheusQueryV2Builder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public PrometheusQueryV2Builder range(Boolean range) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).range = range;
        return this;
    }
    
    public PrometheusQueryV2Builder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public PrometheusQueryV2Builder resultAssertions(com.grafana.foundation.cog.Builder<ResultAssertions> resultAssertions) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
    ResultAssertions resultAssertionsResource = resultAssertions.build();
        ((Dataquery) this.internal.spec).resultAssertions = resultAssertionsResource;
        return this;
    }
    
    public PrometheusQueryV2Builder scopes(List<com.grafana.foundation.cog.Builder<Scopes>> scopes) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        List<Scopes> scopesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<Scopes> r1 : scopes) {
                Scopes scopesDepth1 = r1.build();
                scopesResources.add(scopesDepth1); 
        }
        ((Dataquery) this.internal.spec).scopes = scopesResources;
        return this;
    }
    
    public PrometheusQueryV2Builder timeRange(com.grafana.foundation.cog.Builder<TimeRange> timeRange) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
    TimeRange timeRangeResource = timeRange.build();
        ((Dataquery) this.internal.spec).timeRange = timeRangeResource;
        return this;
    }
    
    public PrometheusQueryV2Builder interval(String interval) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).interval = interval;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
