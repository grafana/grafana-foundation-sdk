// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bigquery;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;

public class BigqueryQueryBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public BigqueryQueryBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "grafana-bigquery-datasource";
    }
    public BigqueryQueryBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public BigqueryQueryBuilder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public BigqueryQueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public BigqueryQueryBuilder dataset(String dataset) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).dataset = dataset;
        return this;
    }
    
    public BigqueryQueryBuilder table(String table) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).table = table;
        return this;
    }
    
    public BigqueryQueryBuilder project(String project) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).project = project;
        return this;
    }
    
    public BigqueryQueryBuilder format(QueryFormat format) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).format = format;
        return this;
    }
    
    public BigqueryQueryBuilder rawQuery(Boolean rawQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).rawQuery = rawQuery;
        return this;
    }
    
    public BigqueryQueryBuilder rawSql(String rawSql) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).rawSql = rawSql;
        return this;
    }
    
    public BigqueryQueryBuilder location(String location) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).location = location;
        return this;
    }
    
    public BigqueryQueryBuilder partitioned(Boolean partitioned) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).partitioned = partitioned;
        return this;
    }
    
    public BigqueryQueryBuilder partitionedField(String partitionedField) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).partitionedField = partitionedField;
        return this;
    }
    
    public BigqueryQueryBuilder convertToUTC(Boolean convertToUTC) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).convertToUTC = convertToUTC;
        return this;
    }
    
    public BigqueryQueryBuilder sharded(Boolean sharded) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).sharded = sharded;
        return this;
    }
    
    public BigqueryQueryBuilder queryPriority(QueryPriority queryPriority) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryPriority = queryPriority;
        return this;
    }
    
    public BigqueryQueryBuilder timeShift(String timeShift) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).timeShift = timeShift;
        return this;
    }
    
    public BigqueryQueryBuilder editorMode(EditorMode editorMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).editorMode = editorMode;
        return this;
    }
    
    public BigqueryQueryBuilder sql(com.grafana.foundation.cog.Builder<SQLExpression> sql) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
    SQLExpression sqlResource = sql.build();
        ((Dataquery) this.internal.spec).sql = sqlResource;
        return this;
    }
    
    public BigqueryQueryBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public BigqueryQueryBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public BigqueryQueryBuilder queryType(String queryType) {
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
