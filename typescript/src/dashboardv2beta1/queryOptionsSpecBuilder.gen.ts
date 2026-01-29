// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class QueryOptionsSpecBuilder implements cog.Builder<dashboardv2beta1.QueryOptionsSpec> {
    protected readonly internal: dashboardv2beta1.QueryOptionsSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultQueryOptionsSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.QueryOptionsSpec {
        return this.internal;
    }

    timeFrom(timeFrom: string): this {
        this.internal.timeFrom = timeFrom;
        return this;
    }

    maxDataPoints(maxDataPoints: number): this {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }

    timeShift(timeShift: string): this {
        this.internal.timeShift = timeShift;
        return this;
    }

    queryCachingTTL(queryCachingTTL: number): this {
        this.internal.queryCachingTTL = queryCachingTTL;
        return this;
    }

    interval(interval: string): this {
        this.internal.interval = interval;
        return this;
    }

    cacheTimeout(cacheTimeout: string): this {
        this.internal.cacheTimeout = cacheTimeout;
        return this;
    }

    hideTimeOverride(hideTimeOverride: boolean): this {
        this.internal.hideTimeOverride = hideTimeOverride;
        return this;
    }

    timeCompare(timeCompare: string): this {
        this.internal.timeCompare = timeCompare;
        return this;
    }
}

