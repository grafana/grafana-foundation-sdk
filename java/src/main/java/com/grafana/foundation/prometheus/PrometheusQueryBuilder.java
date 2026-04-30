// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.prometheus;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;

public class PrometheusQueryBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public PrometheusQueryBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "prometheus";
    }
    public PrometheusQueryBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public PrometheusQueryBuilder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public PrometheusQueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public PrometheusQueryBuilder expr(String expr) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).expr = expr;
        return this;
    }
    
    public PrometheusQueryBuilder instant() {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).instant = true;
        ((Dataquery) this.internal.spec).range = false;
        return this;
    }
    
    public PrometheusQueryBuilder range() {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).range = true;
        ((Dataquery) this.internal.spec).instant = false;
        return this;
    }
    
    public PrometheusQueryBuilder exemplar(Boolean exemplar) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).exemplar = exemplar;
        return this;
    }
    
    public PrometheusQueryBuilder editorMode(QueryEditorMode editorMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).editorMode = editorMode;
        return this;
    }
    
    public PrometheusQueryBuilder format(PromQueryFormat format) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).format = format;
        return this;
    }
    
    public PrometheusQueryBuilder legendFormat(String legendFormat) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).legendFormat = legendFormat;
        return this;
    }
    
    public PrometheusQueryBuilder intervalFactor(Double intervalFactor) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).intervalFactor = intervalFactor;
        return this;
    }
    
    public PrometheusQueryBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public PrometheusQueryBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public PrometheusQueryBuilder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public PrometheusQueryBuilder interval(String interval) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).interval = interval;
        return this;
    }
    
    public PrometheusQueryBuilder rangeAndInstant() {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).range = true;
        ((Dataquery) this.internal.spec).instant = true;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
