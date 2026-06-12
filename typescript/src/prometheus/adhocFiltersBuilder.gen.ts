// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as prometheus from '../prometheus';

export class AdhocFiltersBuilder implements cog.Builder<prometheus.AdhocFilters> {
    protected readonly internal: prometheus.AdhocFilters;

    constructor() {
        this.internal = prometheus.defaultAdhocFilters();
    }

    /**
     * Builds the object.
     */
    build(): prometheus.AdhocFilters {
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

