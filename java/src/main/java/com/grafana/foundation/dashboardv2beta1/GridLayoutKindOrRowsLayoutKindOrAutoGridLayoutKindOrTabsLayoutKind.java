// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindDeserializer.class)
public class GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected GridLayoutKind gridLayoutKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected RowsLayoutKind rowsLayoutKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected AutoGridLayoutKind autoGridLayoutKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected TabsLayoutKind tabsLayoutKind;
    protected GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind() {}
    public static GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind createGridLayoutKind(GridLayoutKind gridLayoutKind) {
        GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind = new GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind();
        gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind.gridLayoutKind = gridLayoutKind;
        return gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind;
    }
    public static GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind createRowsLayoutKind(RowsLayoutKind rowsLayoutKind) {
        GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind = new GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind();
        gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind.rowsLayoutKind = rowsLayoutKind;
        return gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind;
    }
    public static GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind createAutoGridLayoutKind(AutoGridLayoutKind autoGridLayoutKind) {
        GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind = new GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind();
        gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind.autoGridLayoutKind = autoGridLayoutKind;
        return gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind;
    }
    public static GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind createTabsLayoutKind(TabsLayoutKind tabsLayoutKind) {
        GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind = new GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind();
        gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind.tabsLayoutKind = tabsLayoutKind;
        return gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
