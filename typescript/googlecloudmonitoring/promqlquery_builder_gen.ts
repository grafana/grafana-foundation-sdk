// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as googlecloudmonitoring from '../googlecloudmonitoring';

// PromQL sub-query properties.
export class PromQLQueryBuilder implements cog.Builder<googlecloudmonitoring.PromQLQuery> {
    private readonly internal: googlecloudmonitoring.PromQLQuery;

    constructor() {
        this.internal = googlecloudmonitoring.defaultPromQLQuery();
    }

    build(): googlecloudmonitoring.PromQLQuery {
        return this.internal;
    }

    // GCP project to execute the query against.
    projectName(projectName: string): this {
        this.internal.projectName = projectName;
        return this;
    }

    // PromQL expression/query to be executed.
    expr(expr: string): this {
        this.internal.expr = expr;
        return this;
    }

    // PromQL min step
    step(step: string): this {
        this.internal.step = step;
        return this;
    }
}
