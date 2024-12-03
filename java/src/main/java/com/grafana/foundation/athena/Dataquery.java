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
    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
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
    public String dataqueryName() {
        return "grafana-athena-datasource";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
        protected final Dataquery internal;
        
        public Builder() {
            this.internal = new Dataquery();
        this.rawSQL("");
        }
    public Builder format(FormatOptions format) {
    this.internal.format = format;
        return this;
    }
    
    public Builder connectionArgs(com.grafana.foundation.cog.Builder<ConnectionArgs> connectionArgs) {
    this.internal.connectionArgs = connectionArgs.build();
        return this;
    }
    
    public Builder table(String table) {
    this.internal.table = table;
        return this;
    }
    
    public Builder column(String column) {
    this.internal.column = column;
        return this;
    }
    
    public Builder queryID(String queryID) {
    this.internal.queryID = queryID;
        return this;
    }
    
    public Builder refId(String refId) {
    this.internal.refId = refId;
        return this;
    }
    
    public Builder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    
    public Builder queryType(String queryType) {
    this.internal.queryType = queryType;
        return this;
    }
    
    public Builder rawSQL(String rawSQL) {
    this.internal.rawSQL = rawSQL;
        return this;
    }
    
    public Builder datasource(DataSourceRef datasource) {
    this.internal.datasource = datasource;
        return this;
    }
    public Dataquery build() {
            return this.internal;
        }
    }
}
