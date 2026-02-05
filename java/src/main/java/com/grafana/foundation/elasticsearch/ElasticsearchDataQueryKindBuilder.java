// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;
import java.util.List;

public class ElasticsearchDataQueryKindBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public ElasticsearchDataQueryKindBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "elasticsearch";
    }
    public ElasticsearchDataQueryKindBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public ElasticsearchDataQueryKindBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public ElasticsearchDataQueryKindBuilder alias(String alias) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).alias = alias;
        return this;
    }
    
    public ElasticsearchDataQueryKindBuilder query(String query) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).query = query;
        return this;
    }
    
    public ElasticsearchDataQueryKindBuilder timeField(String timeField) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).timeField = timeField;
        return this;
    }
    
    public ElasticsearchDataQueryKindBuilder bucketAggs(List<BucketAggregation> bucketAggs) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).bucketAggs = bucketAggs;
        return this;
    }
    
    public ElasticsearchDataQueryKindBuilder metrics(List<MetricAggregation> metrics) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).metrics = metrics;
        return this;
    }
    
    public ElasticsearchDataQueryKindBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public ElasticsearchDataQueryKindBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public ElasticsearchDataQueryKindBuilder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
