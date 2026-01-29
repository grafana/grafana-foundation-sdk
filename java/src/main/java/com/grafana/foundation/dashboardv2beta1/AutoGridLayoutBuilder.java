// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.LinkedList;

public class AutoGridLayoutBuilder implements com.grafana.foundation.cog.Builder<AutoGridLayoutKind> {
    protected final AutoGridLayoutKind internal;
    
    public AutoGridLayoutBuilder() {
        this.internal = new AutoGridLayoutKind();
        this.internal.kind = "AutoGridLayout";
    }
    public AutoGridLayoutBuilder maxColumnCount(Double maxColumnCount) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutSpecBuilder().build();
		}
        this.internal.spec.maxColumnCount = maxColumnCount;
        return this;
    }
    
    public AutoGridLayoutBuilder columnWidthMode(AutoGridLayoutSpecColumnWidthMode columnWidthMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutSpecBuilder().build();
		}
        this.internal.spec.columnWidthMode = columnWidthMode;
        return this;
    }
    
    public AutoGridLayoutBuilder columnWidth(Double columnWidth) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutSpecBuilder().build();
		}
        this.internal.spec.columnWidth = columnWidth;
        return this;
    }
    
    public AutoGridLayoutBuilder rowHeightMode(AutoGridLayoutSpecRowHeightMode rowHeightMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutSpecBuilder().build();
		}
        this.internal.spec.rowHeightMode = rowHeightMode;
        return this;
    }
    
    public AutoGridLayoutBuilder rowHeight(Double rowHeight) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutSpecBuilder().build();
		}
        this.internal.spec.rowHeight = rowHeight;
        return this;
    }
    
    public AutoGridLayoutBuilder fillScreen(Boolean fillScreen) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutSpecBuilder().build();
		}
        this.internal.spec.fillScreen = fillScreen;
        return this;
    }
    
    public AutoGridLayoutBuilder items(List<com.grafana.foundation.cog.Builder<AutoGridLayoutItemKind>> items) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutSpecBuilder().build();
		}
        List<AutoGridLayoutItemKind> itemsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<AutoGridLayoutItemKind> r1 : items) {
                AutoGridLayoutItemKind itemsDepth1 = r1.build();
                itemsResources.add(itemsDepth1); 
        }
        this.internal.spec.items = itemsResources;
        return this;
    }
    
    public AutoGridLayoutBuilder item(com.grafana.foundation.cog.Builder<AutoGridLayoutItemKind> item) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutSpecBuilder().build();
		}
		if (this.internal.spec.items == null) {
			this.internal.spec.items = new LinkedList<>();
		}
    AutoGridLayoutItemKind itemResource = item.build();
        this.internal.spec.items.add(itemResource);
        return this;
    }
    public AutoGridLayoutKind build() {
        return this.internal;
    }
}
