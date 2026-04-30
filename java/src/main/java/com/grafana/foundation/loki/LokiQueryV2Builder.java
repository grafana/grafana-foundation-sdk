// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.loki;

import com.grafana.foundation.dashboardv2.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2.Dashboardv2DataQueryKindDatasource;

public class LokiQueryV2Builder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public LokiQueryV2Builder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "loki";
    }
    public LokiQueryV2Builder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public LokiQueryV2Builder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public LokiQueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource) {
    Dashboardv2DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public LokiQueryV2Builder expr(String expr) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).expr = expr;
        return this;
    }
    
    public LokiQueryV2Builder legendFormat(String legendFormat) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).legendFormat = legendFormat;
        return this;
    }
    
    public LokiQueryV2Builder maxLines(Long maxLines) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).maxLines = maxLines;
        return this;
    }
    
    public LokiQueryV2Builder resolution(Long resolution) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).resolution = resolution;
        return this;
    }
    
    public LokiQueryV2Builder editorMode(QueryEditorMode editorMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).editorMode = editorMode;
        return this;
    }
    
    public LokiQueryV2Builder range(Boolean range) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).range = range;
        return this;
    }
    
    public LokiQueryV2Builder instant(Boolean instant) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).instant = instant;
        return this;
    }
    
    public LokiQueryV2Builder step(String step) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).step = step;
        return this;
    }
    
    public LokiQueryV2Builder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public LokiQueryV2Builder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public LokiQueryV2Builder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public LokiQueryV2Builder direction(LokiQueryDirection direction) {
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
