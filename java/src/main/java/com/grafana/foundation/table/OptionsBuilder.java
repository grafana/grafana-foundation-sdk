// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.table;

import java.util.List;
import com.grafana.foundation.common.TableSortByFieldState;
import java.util.LinkedList;
import com.grafana.foundation.common.TableFooterOptions;
import com.grafana.foundation.common.TableCellHeight;

public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder frameIndex(Double frameIndex) {
        this.internal.frameIndex = frameIndex;
        return this;
    }
    
    public OptionsBuilder showHeader(Boolean showHeader) {
        this.internal.showHeader = showHeader;
        return this;
    }
    
    public OptionsBuilder showTypeIcons(Boolean showTypeIcons) {
        this.internal.showTypeIcons = showTypeIcons;
        return this;
    }
    
    public OptionsBuilder sortBy(List<com.grafana.foundation.cog.Builder<TableSortByFieldState>> sortBy) {
        List<TableSortByFieldState> sortByResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<TableSortByFieldState> r1 : sortBy) {
                TableSortByFieldState sortByDepth1 = r1.build();
                sortByResources.add(sortByDepth1); 
        }
        this.internal.sortBy = sortByResources;
        return this;
    }
    
    public OptionsBuilder footer(com.grafana.foundation.cog.Builder<TableFooterOptions> footer) {
    TableFooterOptions footerResource = footer.build();
        this.internal.footer = footerResource;
        return this;
    }
    
    public OptionsBuilder cellHeight(TableCellHeight cellHeight) {
        this.internal.cellHeight = cellHeight;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
