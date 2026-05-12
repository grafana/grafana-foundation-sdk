// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.grafana.foundation.dashboardv2.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2.Dashboardv2DataQueryKindDatasource;
import java.util.List;

public class ElasticsearchQueryV2Builder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public ElasticsearchQueryV2Builder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "elasticsearch";
    }
    public ElasticsearchQueryV2Builder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public ElasticsearchQueryV2Builder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public ElasticsearchQueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource) {
    Dashboardv2DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public ElasticsearchQueryV2Builder alias(String alias) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).alias = alias;
        return this;
    }
    
    public ElasticsearchQueryV2Builder query(String query) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).query = query;
        return this;
    }
    
    public ElasticsearchQueryV2Builder timeField(String timeField) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).timeField = timeField;
        return this;
    }
    
    public ElasticsearchQueryV2Builder bucketAggs(List<BucketAggregation> bucketAggs) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).bucketAggs = bucketAggs;
        return this;
    }
    
    public ElasticsearchQueryV2Builder metrics(List<MetricAggregation> metrics) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).metrics = metrics;
        return this;
    }
    
    public ElasticsearchQueryV2Builder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public ElasticsearchQueryV2Builder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.elasticsearch.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public ElasticsearchQueryV2Builder queryType(String queryType) {
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
