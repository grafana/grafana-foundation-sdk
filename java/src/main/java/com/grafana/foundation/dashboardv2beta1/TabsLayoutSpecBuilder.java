// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.LinkedList;

public class TabsLayoutSpecBuilder implements com.grafana.foundation.cog.Builder<TabsLayoutSpec> {
    protected final TabsLayoutSpec internal;
    
    public TabsLayoutSpecBuilder() {
        this.internal = new TabsLayoutSpec();
    }
    public TabsLayoutSpecBuilder tabs(List<com.grafana.foundation.cog.Builder<TabsLayoutTabKind>> tabs) {
        List<TabsLayoutTabKind> tabsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<TabsLayoutTabKind> r1 : tabs) {
                TabsLayoutTabKind tabsDepth1 = r1.build();
                tabsResources.add(tabsDepth1); 
        }
        this.internal.tabs = tabsResources;
        return this;
    }
    public TabsLayoutSpec build() {
        return this.internal;
    }
}
