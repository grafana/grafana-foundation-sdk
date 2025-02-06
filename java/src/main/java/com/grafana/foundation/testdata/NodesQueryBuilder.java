// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;


public class NodesQueryBuilder implements com.grafana.foundation.cog.Builder<NodesQuery> {
    protected final NodesQuery internal;
    
    public NodesQueryBuilder() {
        this.internal = new NodesQuery();
    }
    public NodesQueryBuilder count(Long count) {
        this.internal.count = count;
        return this;
    }
    
    public NodesQueryBuilder seed(Long seed) {
        this.internal.seed = seed;
        return this;
    }
    
    public NodesQueryBuilder type(NodesQueryType type) {
        this.internal.type = type;
        return this;
    }
    public NodesQuery build() {
        return this.internal;
    }
}
