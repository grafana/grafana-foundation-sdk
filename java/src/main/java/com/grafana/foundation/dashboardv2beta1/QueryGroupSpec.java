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
import com.fasterxml.jackson.annotation.JsonInclude;

public class QueryGroupSpec {
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("queries")
    public List<PanelQueryKind> queries;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("transformations")
    public List<TransformationKind> transformations;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("queryOptions")
    public QueryOptionsSpec queryOptions;
    public QueryGroupSpec() {
        this.queries = new LinkedList<>();
        this.transformations = new LinkedList<>();
        this.queryOptions = new com.grafana.foundation.dashboardv2beta1.QueryOptionsSpec();
    }
    public QueryGroupSpec(List<PanelQueryKind> queries,List<TransformationKind> transformations,QueryOptionsSpec queryOptions) {
        this.queries = queries;
        this.transformations = transformations;
        this.queryOptions = queryOptions;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
