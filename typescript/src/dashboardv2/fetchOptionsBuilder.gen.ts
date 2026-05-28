// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class FetchOptionsBuilder implements cog.Builder<dashboardv2.FetchOptions> {
    protected readonly internal: dashboardv2.FetchOptions;

    constructor() {
        this.internal = dashboardv2.defaultFetchOptions();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.FetchOptions {
        return this.internal;
    }

    method(method: dashboardv2.HttpRequestMethod): this {
        this.internal.method = method;
        return this;
    }

    url(url: string): this {
        this.internal.url = url;
        return this;
    }

    body(body: string): this {
        this.internal.body = body;
        return this;
    }

    // These are 2D arrays of strings, each representing a key-value pair
    // We are defining them this way because we can't generate a go struct that
    // that would have exactly two strings in each sub-array
    queryParams(queryParams: string[][]): this {
        this.internal.queryParams = queryParams;
        return this;
    }

    headers(headers: string[][]): this {
        this.internal.headers = headers;
        return this;
    }
}

