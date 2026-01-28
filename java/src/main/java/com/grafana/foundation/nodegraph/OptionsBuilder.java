// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.nodegraph;


public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder nodes(com.grafana.foundation.cog.Builder<NodeOptions> nodes) {
    NodeOptions nodesResource = nodes.build();
        this.internal.nodes = nodesResource;
        return this;
    }
    
    public OptionsBuilder edges(com.grafana.foundation.cog.Builder<EdgeOptions> edges) {
    EdgeOptions edgesResource = edges.build();
        this.internal.edges = edgesResource;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
