// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.LinkedList;

public class TabsBuilder implements com.grafana.foundation.cog.Builder<TabsLayoutKind> {
    protected final TabsLayoutKind internal;
    
    public TabsBuilder() {
        this.internal = new TabsLayoutKind();
        this.internal.kind = "TabsLayout";
    }
    public TabsBuilder tabs(List<com.grafana.foundation.cog.Builder<TabsLayoutTabKind>> tabs) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TabsLayoutSpec();
		}
        List<TabsLayoutTabKind> tabsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<TabsLayoutTabKind> r1 : tabs) {
                TabsLayoutTabKind tabsDepth1 = r1.build();
                tabsResources.add(tabsDepth1); 
        }
        this.internal.spec.tabs = tabsResources;
        return this;
    }
    
    public TabsBuilder tab(com.grafana.foundation.cog.Builder<TabsLayoutTabKind> tab) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TabsLayoutSpec();
		}
		if (this.internal.spec.tabs == null) {
			this.internal.spec.tabs = new LinkedList<>();
		}
    TabsLayoutTabKind tabResource = tab.build();
        this.internal.spec.tabs.add(tabResource);
        return this;
    }
    public TabsLayoutKind build() {
        return this.internal;
    }
}
