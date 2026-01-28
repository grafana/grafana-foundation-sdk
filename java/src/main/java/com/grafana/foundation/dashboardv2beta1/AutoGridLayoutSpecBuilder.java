// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.LinkedList;

public class AutoGridLayoutSpecBuilder implements com.grafana.foundation.cog.Builder<AutoGridLayoutSpec> {
    protected final AutoGridLayoutSpec internal;
    
    public AutoGridLayoutSpecBuilder() {
        this.internal = new AutoGridLayoutSpec();
    }
    public AutoGridLayoutSpecBuilder maxColumnCount(Double maxColumnCount) {
        this.internal.maxColumnCount = maxColumnCount;
        return this;
    }
    
    public AutoGridLayoutSpecBuilder columnWidthMode(AutoGridLayoutSpecColumnWidthMode columnWidthMode) {
        this.internal.columnWidthMode = columnWidthMode;
        return this;
    }
    
    public AutoGridLayoutSpecBuilder columnWidth(Double columnWidth) {
        this.internal.columnWidth = columnWidth;
        return this;
    }
    
    public AutoGridLayoutSpecBuilder rowHeightMode(AutoGridLayoutSpecRowHeightMode rowHeightMode) {
        this.internal.rowHeightMode = rowHeightMode;
        return this;
    }
    
    public AutoGridLayoutSpecBuilder rowHeight(Double rowHeight) {
        this.internal.rowHeight = rowHeight;
        return this;
    }
    
    public AutoGridLayoutSpecBuilder fillScreen(Boolean fillScreen) {
        this.internal.fillScreen = fillScreen;
        return this;
    }
    
    public AutoGridLayoutSpecBuilder items(List<com.grafana.foundation.cog.Builder<AutoGridLayoutItemKind>> items) {
        List<AutoGridLayoutItemKind> itemsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<AutoGridLayoutItemKind> r1 : items) {
                AutoGridLayoutItemKind itemsDepth1 = r1.build();
                itemsResources.add(itemsDepth1); 
        }
        this.internal.items = itemsResources;
        return this;
    }
    public AutoGridLayoutSpec build() {
        return this.internal;
    }
}
