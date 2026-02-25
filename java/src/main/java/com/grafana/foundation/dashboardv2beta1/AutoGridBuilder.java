// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.LinkedList;

public class AutoGridBuilder implements com.grafana.foundation.cog.Builder<AutoGridLayoutKind> {
    protected final AutoGridLayoutKind internal;
    
    public AutoGridBuilder() {
        this.internal = new AutoGridLayoutKind();
        this.internal.kind = "AutoGridLayout";
    }
    public AutoGridBuilder maxColumnCount(Double maxColumnCount) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutSpec();
		}
        this.internal.spec.maxColumnCount = maxColumnCount;
        return this;
    }
    
    public AutoGridBuilder columnWidthMode(AutoGridLayoutSpecColumnWidthMode columnWidthMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutSpec();
		}
        this.internal.spec.columnWidthMode = columnWidthMode;
        return this;
    }
    
    public AutoGridBuilder columnWidth(Double columnWidth) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutSpec();
		}
        this.internal.spec.columnWidth = columnWidth;
        return this;
    }
    
    public AutoGridBuilder rowHeightMode(AutoGridLayoutSpecRowHeightMode rowHeightMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutSpec();
		}
        this.internal.spec.rowHeightMode = rowHeightMode;
        return this;
    }
    
    public AutoGridBuilder rowHeight(Double rowHeight) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutSpec();
		}
        this.internal.spec.rowHeight = rowHeight;
        return this;
    }
    
    public AutoGridBuilder fillScreen(Boolean fillScreen) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutSpec();
		}
        this.internal.spec.fillScreen = fillScreen;
        return this;
    }
    
    public AutoGridBuilder items(List<com.grafana.foundation.cog.Builder<AutoGridLayoutItemKind>> items) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutSpec();
		}
        List<AutoGridLayoutItemKind> itemsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<AutoGridLayoutItemKind> r1 : items) {
                AutoGridLayoutItemKind itemsDepth1 = r1.build();
                itemsResources.add(itemsDepth1); 
        }
        this.internal.spec.items = itemsResources;
        return this;
    }
    
    public AutoGridBuilder item(com.grafana.foundation.cog.Builder<AutoGridLayoutItemKind> item) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutSpec();
		}
		if (this.internal.spec.items == null) {
			this.internal.spec.items = new LinkedList<>();
		}
    AutoGridLayoutItemKind itemResource = item.build();
        this.internal.spec.items.add(itemResource);
        return this;
    }
    
    public AutoGridBuilder withItem(String name) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutSpec();
		}
		if (this.internal.spec.items == null) {
			this.internal.spec.items = new LinkedList<>();
		}
    ElementReference elementReference = new ElementReference();
        elementReference.kind = "ElementReference";
        elementReference.name = name;
    AutoGridLayoutItemSpec autoGridLayoutItemSpec = new AutoGridLayoutItemSpec();
        autoGridLayoutItemSpec.element = elementReference;
    AutoGridLayoutItemKind autoGridLayoutItemKind = new AutoGridLayoutItemKind();
        autoGridLayoutItemKind.kind = "AutoGridLayoutItem";
        autoGridLayoutItemKind.spec = autoGridLayoutItemSpec;
        this.internal.spec.items.add(autoGridLayoutItemKind);
        return this;
    }
    public AutoGridLayoutKind build() {
        return this.internal;
    }
}
