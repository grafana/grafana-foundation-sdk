// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.cog.variants.Dataquery;

@JsonDeserialize(using = QueryDeserializer.class)
public class Query {
    // Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasourceUid")
    public String datasourceUid;
    // JSON is the raw JSON query and includes the above properties as well as custom properties.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("model")
    public Dataquery model;
    // QueryType is an optional identifier for the type of query.
    // It can be used to distinguish different types of queries.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("queryType")
    public String queryType;
    // RefID is the unique identifier of the query, set by the frontend call.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("refId")
    public String refId;
    // RelativeTimeRange is the per query start and end time
    // for requests.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("relativeTimeRange")
    public RelativeTimeRange relativeTimeRange;
    public Query() {
    }
    public Query(String datasourceUid,Dataquery model,String queryType,String refId,RelativeTimeRange relativeTimeRange) {
        this.datasourceUid = datasourceUid;
        this.model = model;
        this.queryType = queryType;
        this.refId = refId;
        this.relativeTimeRange = relativeTimeRange;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
