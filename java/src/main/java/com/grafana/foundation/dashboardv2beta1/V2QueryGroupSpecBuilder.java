// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.LinkedList;

public class V2QueryGroupSpecBuilder implements com.grafana.foundation.cog.Builder<V2QueryGroupSpec> {
    protected final V2QueryGroupSpec internal;
    
    public V2QueryGroupSpecBuilder() {
        this.internal = new V2QueryGroupSpec();
    }
    public V2QueryGroupSpecBuilder queries(List<com.grafana.foundation.cog.Builder<PanelQueryKind>> queries) {
        List<PanelQueryKind> queriesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<PanelQueryKind> r1 : queries) {
                PanelQueryKind queriesDepth1 = r1.build();
                queriesResources.add(queriesDepth1); 
        }
        this.internal.queries = queriesResources;
        return this;
    }
    
    public V2QueryGroupSpecBuilder transformations(List<com.grafana.foundation.cog.Builder<V2TransformationKind>> transformations) {
        List<V2TransformationKind> transformationsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<V2TransformationKind> r1 : transformations) {
                V2TransformationKind transformationsDepth1 = r1.build();
                transformationsResources.add(transformationsDepth1); 
        }
        this.internal.transformations = transformationsResources;
        return this;
    }
    
    public V2QueryGroupSpecBuilder queryOptions(com.grafana.foundation.cog.Builder<QueryOptionsSpec> queryOptions) {
    QueryOptionsSpec queryOptionsResource = queryOptions.build();
        this.internal.queryOptions = queryOptionsResource;
        return this;
    }
    public V2QueryGroupSpec build() {
        return this.internal;
    }
}
