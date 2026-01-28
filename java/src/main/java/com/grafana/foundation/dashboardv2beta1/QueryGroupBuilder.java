// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.LinkedList;

public class QueryGroupBuilder implements com.grafana.foundation.cog.Builder<QueryGroupKind> {
    protected final QueryGroupKind internal;
    
    public QueryGroupBuilder() {
        this.internal = new QueryGroupKind();
        this.internal.kind = "QueryGroup";
    }
    public QueryGroupBuilder targets(List<com.grafana.foundation.cog.Builder<PanelQueryKind>> queries) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryGroupSpec();
		}
        List<PanelQueryKind> queriesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<PanelQueryKind> r1 : queries) {
                PanelQueryKind queriesDepth1 = r1.build();
                queriesResources.add(queriesDepth1); 
        }
        this.internal.spec.queries = queriesResources;
        return this;
    }
    
    public QueryGroupBuilder target(com.grafana.foundation.cog.Builder<PanelQueryKind> querie) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryGroupSpec();
		}
		if (this.internal.spec.queries == null) {
			this.internal.spec.queries = new LinkedList<>();
		}
    PanelQueryKind querieResource = querie.build();
        this.internal.spec.queries.add(querieResource);
        return this;
    }
    
    public QueryGroupBuilder transformations(List<com.grafana.foundation.cog.Builder<TransformationKind>> transformations) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryGroupSpec();
		}
        List<TransformationKind> transformationsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<TransformationKind> r1 : transformations) {
                TransformationKind transformationsDepth1 = r1.build();
                transformationsResources.add(transformationsDepth1); 
        }
        this.internal.spec.transformations = transformationsResources;
        return this;
    }
    
    public QueryGroupBuilder transformation(com.grafana.foundation.cog.Builder<TransformationKind> transformation) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryGroupSpec();
		}
		if (this.internal.spec.transformations == null) {
			this.internal.spec.transformations = new LinkedList<>();
		}
    TransformationKind transformationResource = transformation.build();
        this.internal.spec.transformations.add(transformationResource);
        return this;
    }
    
    public QueryGroupBuilder queryOptions(com.grafana.foundation.cog.Builder<QueryOptionsSpec> queryOptions) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.QueryGroupSpec();
		}
    QueryOptionsSpec queryOptionsResource = queryOptions.build();
        this.internal.spec.queryOptions = queryOptionsResource;
        return this;
    }
    public QueryGroupKind build() {
        return this.internal;
    }
}
