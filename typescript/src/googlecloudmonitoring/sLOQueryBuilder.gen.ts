// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as googlecloudmonitoring from '../googlecloudmonitoring';

// SLO sub-query properties.
export class SLOQueryBuilder implements cog.Builder<googlecloudmonitoring.SLOQuery> {
    protected readonly internal: googlecloudmonitoring.SLOQuery;

    constructor() {
        this.internal = googlecloudmonitoring.defaultSLOQuery();
    }

    /**
     * Builds the object.
     */
    build(): googlecloudmonitoring.SLOQuery {
        return this.internal;
    }

    // GCP project to execute the query against.
    projectName(projectName: string): this {
        this.internal.projectName = projectName;
        return this;
    }

    // Alignment function to be used. Defaults to ALIGN_MEAN.
    perSeriesAligner(perSeriesAligner: string): this {
        this.internal.perSeriesAligner = perSeriesAligner;
        return this;
    }

    // Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    alignmentPeriod(alignmentPeriod: string): this {
        this.internal.alignmentPeriod = alignmentPeriod;
        return this;
    }

    // SLO selector.
    selectorName(selectorName: string): this {
        this.internal.selectorName = selectorName;
        return this;
    }

    // ID for the service the SLO is in.
    serviceId(serviceId: string): this {
        this.internal.serviceId = serviceId;
        return this;
    }

    // Name for the service the SLO is in.
    serviceName(serviceName: string): this {
        this.internal.serviceName = serviceName;
        return this;
    }

    // ID for the SLO.
    sloId(sloId: string): this {
        this.internal.sloId = sloId;
        return this;
    }

    // Name of the SLO.
    sloName(sloName: string): this {
        this.internal.sloName = sloName;
        return this;
    }

    // SLO goal value.
    goal(goal: number): this {
        this.internal.goal = goal;
        return this;
    }

    // Specific lookback period for the SLO.
    lookbackPeriod(lookbackPeriod: string): this {
        this.internal.lookbackPeriod = lookbackPeriod;
        return this;
    }
}
