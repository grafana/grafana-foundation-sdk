// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.nodegraph;

import java.util.List;

public class NodeOptionsBuilder implements com.grafana.foundation.cog.Builder<NodeOptions> {
    protected final NodeOptions internal;
    
    public NodeOptionsBuilder() {
        this.internal = new NodeOptions();
    }
    public NodeOptionsBuilder mainStatUnit(String mainStatUnit) {
    this.internal.mainStatUnit = mainStatUnit;
        return this;
    }
    
    public NodeOptionsBuilder secondaryStatUnit(String secondaryStatUnit) {
    this.internal.secondaryStatUnit = secondaryStatUnit;
        return this;
    }
    
    public NodeOptionsBuilder arcs(com.grafana.foundation.cog.Builder<List<ArcOption>> arcs) {
    this.internal.arcs = arcs.build();
        return this;
    }
    public NodeOptions build() {
        return this.internal;
    }
}
