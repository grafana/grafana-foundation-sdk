// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;


public class NamedRoutingTreeBuilder implements com.grafana.foundation.cog.Builder<NamedRoutingTree> {
    protected final NamedRoutingTree internal;
    
    public NamedRoutingTreeBuilder() {
        this.internal = new NamedRoutingTree();
    }
    public NamedRoutingTreeBuilder routingTree(String routingTree) {
        this.internal.routingTree = routingTree;
        return this;
    }
    public NamedRoutingTree build() {
        return this.internal;
    }
}
