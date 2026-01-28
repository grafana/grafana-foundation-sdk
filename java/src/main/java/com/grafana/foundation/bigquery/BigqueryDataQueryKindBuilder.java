// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bigquery;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;
import com.grafana.foundation.common.DataSourceRef;

public class BigqueryDataQueryKindBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public BigqueryDataQueryKindBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "grafana-bigquery-datasource";
    }
    public BigqueryDataQueryKindBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder dataset(String dataset) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).dataset = dataset;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder table(String table) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).table = table;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder project(String project) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).project = project;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder format(QueryFormat format) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).format = format;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder rawQuery(Boolean rawQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).rawQuery = rawQuery;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder rawSql(String rawSql) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).rawSql = rawSql;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder location(String location) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).location = location;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder partitioned(Boolean partitioned) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).partitioned = partitioned;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder partitionedField(String partitionedField) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).partitionedField = partitionedField;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder convertToUTC(Boolean convertToUTC) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).convertToUTC = convertToUTC;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder sharded(Boolean sharded) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).sharded = sharded;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder queryPriority(QueryPriority queryPriority) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryPriority = queryPriority;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder timeShift(String timeShift) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).timeShift = timeShift;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder editorMode(EditorMode editorMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).editorMode = editorMode;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder sql(com.grafana.foundation.cog.Builder<SQLExpression> sql) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
    SQLExpression sqlResource = sql.build();
        ((Dataquery) this.internal.spec).sql = sqlResource;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public BigqueryDataQueryKindBuilder oldDatasource(DataSourceRef datasource) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.bigquery.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).datasource = datasource;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
