// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class GridLayoutSpec {
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("items")
    public List<GridLayoutItemKind> items;
    public GridLayoutSpec() {
        this.items = new LinkedList<>();
    }
    public GridLayoutSpec(List<GridLayoutItemKind> items) {
        this.items = items;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
