// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.loki;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;

public class LokiQueryBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public LokiQueryBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "loki";
    }
    public LokiQueryBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public LokiQueryBuilder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public LokiQueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public LokiQueryBuilder expr(String expr) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).expr = expr;
        return this;
    }
    
    public LokiQueryBuilder legendFormat(String legendFormat) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).legendFormat = legendFormat;
        return this;
    }
    
    public LokiQueryBuilder maxLines(Long maxLines) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).maxLines = maxLines;
        return this;
    }
    
    public LokiQueryBuilder resolution(Long resolution) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).resolution = resolution;
        return this;
    }
    
    public LokiQueryBuilder editorMode(QueryEditorMode editorMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).editorMode = editorMode;
        return this;
    }
    
    public LokiQueryBuilder range(Boolean range) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).range = range;
        return this;
    }
    
    public LokiQueryBuilder instant(Boolean instant) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).instant = instant;
        return this;
    }
    
    public LokiQueryBuilder step(String step) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).step = step;
        return this;
    }
    
    public LokiQueryBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public LokiQueryBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public LokiQueryBuilder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public LokiQueryBuilder direction(LokiQueryDirection direction) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).direction = direction;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
