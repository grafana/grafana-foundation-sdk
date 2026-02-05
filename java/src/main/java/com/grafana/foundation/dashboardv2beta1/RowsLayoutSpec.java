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

public class RowsLayoutSpec {
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("rows")
    public List<RowsLayoutRowKind> rows;
    public RowsLayoutSpec() {
        this.rows = new LinkedList<>();
    }
    public RowsLayoutSpec(List<RowsLayoutRowKind> rows) {
        this.rows = rows;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
