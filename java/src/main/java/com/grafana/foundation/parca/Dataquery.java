// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.parca;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.dashboard.DataSourceRef;

public class Dataquery implements com.grafana.foundation.cog.variants.Dataquery {
    // Specifies the query label selectors.
    @JsonProperty("labelSelector")
    public String labelSelector;
    // Specifies the type of profile to query.
    @JsonProperty("profileTypeId")
    public String profileTypeId;
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    @JsonProperty("refId")
    public String refId;
    // true if query is disabled (ie should not be returned to the dashboard)
    // Note this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
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
        this.refId = "";
    }
    public Dataquery(String labelSelector,String profileTypeId,String refId,Boolean hide,String queryType,DataSourceRef datasource) {
        this.labelSelector = labelSelector;
        this.profileTypeId = profileTypeId;
        this.refId = refId;
        this.hide = hide;
        this.queryType = queryType;
        this.datasource = datasource;
    }
    public String dataqueryName() {
        return "parca";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
