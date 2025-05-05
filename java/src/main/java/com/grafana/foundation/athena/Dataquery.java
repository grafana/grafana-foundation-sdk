// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.athena;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.dashboard.DataSourceRef;

// Manually converted from https://github.com/grafana/athena-datasource/blob/57ad707147b7a11e9a521a836d6bf9799877e0e3/src/types.ts
public class Dataquery implements com.grafana.foundation.cog.variants.Dataquery {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("format")
    public FormatOptions format;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("connectionArgs")
    public ConnectionArgs connectionArgs;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("table")
    public String table;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("column")
    public String column;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("queryID")
    public String queryID;
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
    @JsonProperty("rawSQL")
    public String rawSQL;
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    public Dataquery() {
        this.format = FormatOptions.TIME_SERIES;
        this.connectionArgs = new com.grafana.foundation.athena.ConnectionArgs();
        this.refId = "";
        this.rawSQL = "";
    }
    public Dataquery(FormatOptions format,ConnectionArgs connectionArgs,String table,String column,String queryID,String refId,Boolean hide,String queryType,String rawSQL,DataSourceRef datasource) {
        this.format = format;
        this.connectionArgs = connectionArgs;
        this.table = table;
        this.column = column;
        this.queryID = queryID;
        this.refId = refId;
        this.hide = hide;
        this.queryType = queryType;
        this.rawSQL = rawSQL;
        this.datasource = datasource;
    }
    public String dataqueryName() {
        return "grafana-athena-datasource";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
