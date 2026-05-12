// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bigquery;

import com.grafana.foundation.dashboardv2.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2.Dashboardv2DataQueryKindDatasource;

public class BigqueryQueryV2Builder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public BigqueryQueryV2Builder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "grafana-bigquery-datasource";
    }
    public BigqueryQueryV2Builder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public BigqueryQueryV2Builder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public BigqueryQueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource) {
    Dashboardv2DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public BigqueryQueryV2Builder dataset(String dataset) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).dataset = dataset;
        return this;
    }
    
    public BigqueryQueryV2Builder table(String table) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).table = table;
        return this;
    }
    
    public BigqueryQueryV2Builder project(String project) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).project = project;
        return this;
    }
    
    public BigqueryQueryV2Builder format(QueryFormat format) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).format = format;
        return this;
    }
    
    public BigqueryQueryV2Builder rawQuery(Boolean rawQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).rawQuery = rawQuery;
        return this;
    }
    
    public BigqueryQueryV2Builder rawSql(String rawSql) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).rawSql = rawSql;
        return this;
    }
    
    public BigqueryQueryV2Builder location(String location) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).location = location;
        return this;
    }
    
    public BigqueryQueryV2Builder partitioned(Boolean partitioned) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).partitioned = partitioned;
        return this;
    }
    
    public BigqueryQueryV2Builder partitionedField(String partitionedField) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).partitionedField = partitionedField;
        return this;
    }
    
    public BigqueryQueryV2Builder convertToUTC(Boolean convertToUTC) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).convertToUTC = convertToUTC;
        return this;
    }
    
    public BigqueryQueryV2Builder sharded(Boolean sharded) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).sharded = sharded;
        return this;
    }
    
    public BigqueryQueryV2Builder queryPriority(QueryPriority queryPriority) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryPriority = queryPriority;
        return this;
    }
    
    public BigqueryQueryV2Builder timeShift(String timeShift) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).timeShift = timeShift;
        return this;
    }
    
    public BigqueryQueryV2Builder editorMode(EditorMode editorMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).editorMode = editorMode;
        return this;
    }
    
    public BigqueryQueryV2Builder sql(com.grafana.foundation.cog.Builder<SQLExpression> sql) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
    SQLExpression sqlResource = sql.build();
        ((Dataquery) this.internal.spec).sql = sqlResource;
        return this;
    }
    
    public BigqueryQueryV2Builder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public BigqueryQueryV2Builder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public BigqueryQueryV2Builder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
