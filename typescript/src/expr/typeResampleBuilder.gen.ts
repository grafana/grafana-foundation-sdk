// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as expr from '../expr';
import * as dashboard from '../dashboard';

export class TypeResampleBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: expr.TypeResample;

    constructor() {
        this.internal = expr.defaultTypeResample();
        this.internal.type = "resample";
    }

    /**
     * Builds the object.
     */
    build(): expr.TypeResample {
        return this.internal;
    }

    // The datasource
    datasource(datasource: dashboard.DataSourceRef): this {
        this.internal.datasource = datasource;
        return this;
    }

    // The downsample function
    // Possible enum values:
    //  - `"sum"` 
    //  - `"mean"` 
    //  - `"min"` 
    //  - `"max"` 
    //  - `"count"` 
    //  - `"last"` 
    //  - `"median"` 
    downsampler(downsampler: "sum" | "mean" | "min" | "max" | "count" | "last" | "median"): this {
        this.internal.downsampler = downsampler;
        return this;
    }

    // The math expression
    expression(expression: string): this {
        if (!(expression.length >= 1)) {
            throw new Error("expression.length must be >= 1");
        }
        this.internal.expression = expression;
        return this;
    }

    // true if query is disabled (ie should not be returned to the dashboard)
    // NOTE: this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }

    // Interval is the suggested duration between time points in a time series query.
    // NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    // from the interval required to fill a pixels in the visualization
    intervalMs(intervalMs: number): this {
        this.internal.intervalMs = intervalMs;
        return this;
    }

    // MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    // NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    // from the number of pixels visible in a visualization
    maxDataPoints(maxDataPoints: number): this {
        this.internal.maxDataPoints = maxDataPoints;
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

    // Optionally define expected query result behavior
    resultAssertions(resultAssertions: {
	// Maximum frame count
	maxFrames?: number;
	// Type asserts that the frame matches a known type structure.
	// Possible enum values:
	//  - `""` 
	//  - `"timeseries-wide"` 
	//  - `"timeseries-long"` 
	//  - `"timeseries-many"` 
	//  - `"timeseries-multi"` 
	//  - `"directory-listing"` 
	//  - `"table"` 
	//  - `"numeric-wide"` 
	//  - `"numeric-multi"` 
	//  - `"numeric-long"` 
	//  - `"log-lines"` 
	type?: "" | "timeseries-wide" | "timeseries-long" | "timeseries-many" | "timeseries-multi" | "directory-listing" | "table" | "numeric-wide" | "numeric-multi" | "numeric-long" | "log-lines";
	// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
	// contract documentation https://grafana.github.io/dataplane/contract/.
	typeVersion: number[];
}): this {
        this.internal.resultAssertions = resultAssertions;
        return this;
    }

    // TimeRange represents the query range
    // NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    // NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    timeRange(timeRange: {
	// From is the start time of the query.
	from: string;
	// To is the end time of the query.
	to: string;
}): this {
        this.internal.timeRange = timeRange;
        return this;
    }

    // The upsample function
    // Possible enum values:
    //  - `"pad"` Use the last seen value
    //  - `"backfilling"` backfill
    //  - `"fillna"` Do not fill values (nill)
    upsampler(upsampler: "pad" | "backfilling" | "fillna"): this {
        this.internal.upsampler = upsampler;
        return this;
    }

    // The time duration
    window(window: string): this {
        if (!(window.length >= 1)) {
            throw new Error("window.length must be >= 1");
        }
        this.internal.window = window;
        return this;
    }
}
