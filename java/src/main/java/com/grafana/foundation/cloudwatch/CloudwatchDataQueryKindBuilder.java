// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;

public class CloudwatchDataQueryKindBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public CloudwatchDataQueryKindBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "cloudwatch";
    }
    public CloudwatchDataQueryKindBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public CloudwatchDataQueryKindBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public CloudwatchDataQueryKindBuilder cloudWatchMetricsQuery(com.grafana.foundation.cog.Builder<CloudWatchMetricsQuery> cloudWatchMetricsQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.cloudwatch.CloudWatchQuery();
		}
    CloudWatchMetricsQuery cloudWatchMetricsQueryResource = cloudWatchMetricsQuery.build();
        ((CloudWatchQuery) this.internal.spec).cloudWatchMetricsQuery = cloudWatchMetricsQueryResource;
        return this;
    }
    
    public CloudwatchDataQueryKindBuilder cloudWatchLogsQuery(com.grafana.foundation.cog.Builder<CloudWatchLogsQuery> cloudWatchLogsQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.cloudwatch.CloudWatchQuery();
		}
    CloudWatchLogsQuery cloudWatchLogsQueryResource = cloudWatchLogsQuery.build();
        ((CloudWatchQuery) this.internal.spec).cloudWatchLogsQuery = cloudWatchLogsQueryResource;
        return this;
    }
    
    public CloudwatchDataQueryKindBuilder cloudWatchAnnotationQuery(com.grafana.foundation.cog.Builder<CloudWatchAnnotationQuery> cloudWatchAnnotationQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.cloudwatch.CloudWatchQuery();
		}
    CloudWatchAnnotationQuery cloudWatchAnnotationQueryResource = cloudWatchAnnotationQuery.build();
        ((CloudWatchQuery) this.internal.spec).cloudWatchAnnotationQuery = cloudWatchAnnotationQueryResource;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
