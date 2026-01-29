// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.LinkedList;

public class GridLayoutBuilder implements com.grafana.foundation.cog.Builder<GridLayoutKind> {
    protected final GridLayoutKind internal;
    
    public GridLayoutBuilder() {
        this.internal = new GridLayoutKind();
        this.internal.kind = "GridLayout";
    }
    public GridLayoutBuilder items(List<com.grafana.foundation.cog.Builder<GridLayoutItemKind>> items) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.GridLayoutSpec();
		}
        List<GridLayoutItemKind> itemsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<GridLayoutItemKind> r1 : items) {
                GridLayoutItemKind itemsDepth1 = r1.build();
                itemsResources.add(itemsDepth1); 
        }
        this.internal.spec.items = itemsResources;
        return this;
    }
    
    public GridLayoutBuilder item(com.grafana.foundation.cog.Builder<GridLayoutItemKind> item) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.GridLayoutSpec();
		}
		if (this.internal.spec.items == null) {
			this.internal.spec.items = new LinkedList<>();
		}
    GridLayoutItemKind itemResource = item.build();
        this.internal.spec.items.add(itemResource);
        return this;
    }
    public GridLayoutKind build() {
        return this.internal;
    }
}
