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

public class TabsLayoutSpec {
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("tabs")
    public List<TabsLayoutTabKind> tabs;
    public TabsLayoutSpec() {
        this.tabs = new LinkedList<>();
    }
    public TabsLayoutSpec(List<TabsLayoutTabKind> tabs) {
        this.tabs = tabs;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
