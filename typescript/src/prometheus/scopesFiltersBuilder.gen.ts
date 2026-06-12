// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as prometheus from '../prometheus';

export class ScopesFiltersBuilder implements cog.Builder<prometheus.ScopesFilters> {
    protected readonly internal: prometheus.ScopesFilters;

    constructor() {
        this.internal = prometheus.defaultScopesFilters();
    }

    /**
     * Builds the object.
     */
    build(): prometheus.ScopesFilters {
        return this.internal;
    }

    key(key: string): this {
        this.internal.key = key;
        return this;
    }

    operator(operator: string): this {
        this.internal.operator = operator;
        return this;
    }

    value(value: string): this {
        this.internal.value = value;
        return this;
    }

    values(values: string[]): this {
        this.internal.values = values;
        return this;
    }
}

