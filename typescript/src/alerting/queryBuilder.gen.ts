// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

export class QueryBuilder implements cog.Builder<alerting.Query> {
    protected readonly internal: alerting.Query;

    constructor(refId: string) {
        this.internal = alerting.defaultQuery();
        this.internal.refId = refId;
    }

    build(): alerting.Query {
        return this.internal;
    }

    // Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.
    datasourceUid(datasourceUid: string): this {
        this.internal.datasourceUid = datasourceUid;
        return this;
    }

    // JSON is the raw JSON query and includes the above properties as well as custom properties.
    model(model: cog.Builder<cog.Dataquery>): this {
        const modelResource = model.build();
        this.internal.model = modelResource;
        return this;
    }

    // QueryType is an optional identifier for the type of query.
    // It can be used to distinguish different types of queries.
    queryType(queryType: string): this {
        this.internal.queryType = queryType;
        return this;
    }

    // RefID is the unique identifier of the query, set by the frontend call.
    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    // RelativeTimeRange is the per query start and end time
    // for requests.
    relativeTimeRange(relativeTimeRange: alerting.RelativeTimeRange): this {
        this.internal.relativeTimeRange = relativeTimeRange;
        return this;
    }
}
