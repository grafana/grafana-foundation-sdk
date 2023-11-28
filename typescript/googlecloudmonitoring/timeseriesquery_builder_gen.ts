// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as googlecloudmonitoring from '../googlecloudmonitoring';

// Time Series sub-query properties.
export class TimeSeriesQueryBuilder implements cog.OptionsBuilder<googlecloudmonitoring.TimeSeriesQuery> {
    private readonly internal: googlecloudmonitoring.TimeSeriesQuery;

    constructor() {
        this.internal = googlecloudmonitoring.defaultTimeSeriesQuery();
    }

    build(): googlecloudmonitoring.TimeSeriesQuery {
        return this.internal;
    }

    // GCP project to execute the query against.
    projectName(projectName: string): this {
        this.internal.projectName = projectName;
        return this;
    }

    // MQL query to be executed.
    query(query: string): this {
        this.internal.query = query;
        return this;
    }

    // To disable the graphPeriod, it should explictly be set to 'disabled'.
    graphPeriod(graphPeriod: "disabled" | string): this {
        this.internal.graphPeriod = graphPeriod;
        return this;
    }
}
