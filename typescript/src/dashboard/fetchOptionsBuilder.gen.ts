// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// Fetch options
export class FetchOptionsBuilder implements cog.Builder<dashboard.FetchOptions> {
    protected readonly internal: dashboard.FetchOptions;

    constructor() {
        this.internal = dashboard.defaultFetchOptions();
    }

    /**
     * Builds the object.
     */
    build(): dashboard.FetchOptions {
        return this.internal;
    }

    method(method: dashboard.HttpRequestMethod): this {
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
    // We are defining this way because we can't generate a go struct that
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

