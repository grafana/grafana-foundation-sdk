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

    datasourceUid(datasourceUid: string): this {
        this.internal.datasourceUid = datasourceUid;
        return this;
    }

    model(model: cog.Builder<cog.Dataquery>): this {
        const modelResource = model.build();
        this.internal.model = modelResource;
        return this;
    }

    queryType(queryType: string): this {
        this.internal.queryType = queryType;
        return this;
    }

    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    relativeTimeRange(relativeTimeRange: alerting.RelativeTimeRange): this {
        this.internal.relativeTimeRange = relativeTimeRange;
        return this;
    }
}
