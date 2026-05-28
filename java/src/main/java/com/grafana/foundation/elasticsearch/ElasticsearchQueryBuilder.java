// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;
import java.util.List;

public class ElasticsearchQueryBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public ElasticsearchQueryBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "elasticsearch";
    }
    public ElasticsearchQueryBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public ElasticsearchQueryBuilder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public ElasticsearchQueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public ElasticsearchQueryBuilder alias(String alias) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).alias = alias;
        return this;
    }
    
    public ElasticsearchQueryBuilder query(String query) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).query = query;
        return this;
    }
    
    public ElasticsearchQueryBuilder timeField(String timeField) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).timeField = timeField;
        return this;
    }
    
    public ElasticsearchQueryBuilder bucketAggs(List<BucketAggregation> bucketAggs) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).bucketAggs = bucketAggs;
        return this;
    }
    
    public ElasticsearchQueryBuilder metrics(List<MetricAggregation> metrics) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).metrics = metrics;
        return this;
    }
    
    public ElasticsearchQueryBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public ElasticsearchQueryBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public ElasticsearchQueryBuilder queryType(String queryType) {
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
