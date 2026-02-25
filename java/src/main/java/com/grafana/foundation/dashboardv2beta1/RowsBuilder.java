// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.LinkedList;

public class RowsBuilder implements com.grafana.foundation.cog.Builder<RowsLayoutKind> {
    protected final RowsLayoutKind internal;
    
    public RowsBuilder() {
        this.internal = new RowsLayoutKind();
        this.internal.kind = "RowsLayout";
    }
    public RowsBuilder rows(List<com.grafana.foundation.cog.Builder<RowsLayoutRowKind>> rows) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutSpec();
		}
        List<RowsLayoutRowKind> rowsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<RowsLayoutRowKind> r1 : rows) {
                RowsLayoutRowKind rowsDepth1 = r1.build();
                rowsResources.add(rowsDepth1); 
        }
        this.internal.spec.rows = rowsResources;
        return this;
    }
    
    public RowsBuilder row(com.grafana.foundation.cog.Builder<RowsLayoutRowKind> row) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutSpec();
		}
		if (this.internal.spec.rows == null) {
			this.internal.spec.rows = new LinkedList<>();
		}
    RowsLayoutRowKind rowResource = row.build();
        this.internal.spec.rows.add(rowResource);
        return this;
    }
    public RowsLayoutKind build() {
        return this.internal;
    }
}
