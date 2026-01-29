// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.grafanapyroscope;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import com.grafana.foundation.common.DataSourceRef;

public class Dataquery implements com.grafana.foundation.cog.variants.Dataquery {
    // Specifies the query label selectors.
    @JsonProperty("labelSelector")
    public String labelSelector;
    // Specifies the query span selectors.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("spanSelector")
    public List<String> spanSelector;
    // Specifies the type of profile to query.
    @JsonProperty("profileTypeId")
    public String profileTypeId;
    // Allows to group the results.
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("groupBy")
    public List<String> groupBy;
    // Sets the maximum number of nodes in the flamegraph.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("maxNodes")
    public Long maxNodes;
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("refId")
    public String refId;
    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    // Specify the query flavor
    // TODO make this required and give it a default
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("queryType")
    public String queryType;
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    public Dataquery() {
        this.labelSelector = "{}";
        this.profileTypeId = "";
        this.groupBy = new LinkedList<>();
    }
    public Dataquery(String labelSelector,List<String> spanSelector,String profileTypeId,List<String> groupBy,Long maxNodes,String refId,Boolean hide,String queryType,DataSourceRef datasource) {
        this.labelSelector = labelSelector;
        this.spanSelector = spanSelector;
        this.profileTypeId = profileTypeId;
        this.groupBy = groupBy;
        this.maxNodes = maxNodes;
        this.refId = refId;
        this.hide = hide;
        this.queryType = queryType;
        this.datasource = datasource;
    }
    public String dataqueryName() {
        return "grafanapyroscope";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
