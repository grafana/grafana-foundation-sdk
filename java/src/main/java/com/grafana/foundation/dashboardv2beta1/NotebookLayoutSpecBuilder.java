// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.LinkedList;

public class NotebookLayoutSpecBuilder implements com.grafana.foundation.cog.Builder<NotebookLayoutSpec> {
    protected final NotebookLayoutSpec internal;
    
    public NotebookLayoutSpecBuilder() {
        this.internal = new NotebookLayoutSpec();
    }
    public NotebookLayoutSpecBuilder cells(List<com.grafana.foundation.cog.Builder<NotebookLayoutItemKind>> cells) {
        List<NotebookLayoutItemKind> cellsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<NotebookLayoutItemKind> r1 : cells) {
                NotebookLayoutItemKind cellsDepth1 = r1.build();
                cellsResources.add(cellsDepth1); 
        }
        this.internal.cells = cellsResources;
        return this;
    }
    public NotebookLayoutSpec build() {
        return this.internal;
    }
}
