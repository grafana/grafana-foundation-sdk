// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.LinkedList;

public class RowsLayoutSpecBuilder implements com.grafana.foundation.cog.Builder<RowsLayoutSpec> {
    protected final RowsLayoutSpec internal;
    
    public RowsLayoutSpecBuilder() {
        this.internal = new RowsLayoutSpec();
    }
    public RowsLayoutSpecBuilder rows(List<com.grafana.foundation.cog.Builder<RowsLayoutRowKind>> rows) {
        List<RowsLayoutRowKind> rowsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<RowsLayoutRowKind> r1 : rows) {
                RowsLayoutRowKind rowsDepth1 = r1.build();
                rowsResources.add(rowsDepth1); 
        }
        this.internal.rows = rowsResources;
        return this;
    }
    public RowsLayoutSpec build() {
        return this.internal;
    }
}
