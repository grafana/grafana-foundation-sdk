// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class AutoGridLayoutSpec {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("maxColumnCount")
    public Double maxColumnCount;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("columnWidthMode")
    public AutoGridLayoutSpecColumnWidthMode columnWidthMode;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("columnWidth")
    public Double columnWidth;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("rowHeightMode")
    public AutoGridLayoutSpecRowHeightMode rowHeightMode;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("rowHeight")
    public Double rowHeight;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fillScreen")
    public Boolean fillScreen;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("items")
    public List<AutoGridLayoutItemKind> items;
    public AutoGridLayoutSpec() {
        this.maxColumnCount = 3.0;
        this.columnWidthMode = AutoGridLayoutSpecColumnWidthMode.STANDARD;
        this.rowHeightMode = AutoGridLayoutSpecRowHeightMode.STANDARD;
        this.fillScreen = false;
        this.items = new LinkedList<>();
    }
    public AutoGridLayoutSpec(Double maxColumnCount,AutoGridLayoutSpecColumnWidthMode columnWidthMode,Double columnWidth,AutoGridLayoutSpecRowHeightMode rowHeightMode,Double rowHeight,Boolean fillScreen,List<AutoGridLayoutItemKind> items) {
        this.maxColumnCount = maxColumnCount;
        this.columnWidthMode = columnWidthMode;
        this.columnWidth = columnWidth;
        this.rowHeightMode = rowHeightMode;
        this.rowHeight = rowHeight;
        this.fillScreen = fillScreen;
        this.items = items;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
