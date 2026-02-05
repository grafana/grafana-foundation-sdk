// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.nodegraph;

import java.util.List;
import java.util.LinkedList;

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
    
    public NodeOptionsBuilder arcs(List<com.grafana.foundation.cog.Builder<ArcOption>> arcs) {
        List<ArcOption> arcsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<ArcOption> r1 : arcs) {
                ArcOption arcsDepth1 = r1.build();
                arcsResources.add(arcsDepth1); 
        }
        this.internal.arcs = arcsResources;
        return this;
    }
    public NodeOptions build() {
        return this.internal;
    }
}
