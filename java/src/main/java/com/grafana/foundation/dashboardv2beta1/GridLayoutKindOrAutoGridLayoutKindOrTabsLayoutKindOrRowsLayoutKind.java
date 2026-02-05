// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKindDeserializer.class)
public class GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected GridLayoutKind gridLayoutKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected AutoGridLayoutKind autoGridLayoutKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected TabsLayoutKind tabsLayoutKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected RowsLayoutKind rowsLayoutKind;
    protected GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind() {}
    public static GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind createGridLayoutKind(GridLayoutKind gridLayoutKind) {
        GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind = new GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind();
        gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind.gridLayoutKind = gridLayoutKind;
        return gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind;
    }
    public static GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind createAutoGridLayoutKind(AutoGridLayoutKind autoGridLayoutKind) {
        GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind = new GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind();
        gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind.autoGridLayoutKind = autoGridLayoutKind;
        return gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind;
    }
    public static GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind createTabsLayoutKind(TabsLayoutKind tabsLayoutKind) {
        GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind = new GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind();
        gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind.tabsLayoutKind = tabsLayoutKind;
        return gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind;
    }
    public static GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind createRowsLayoutKind(RowsLayoutKind rowsLayoutKind) {
        GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind = new GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind();
        gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind.rowsLayoutKind = rowsLayoutKind;
        return gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
