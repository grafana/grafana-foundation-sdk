// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as prometheus from '../prometheus';

export class ScopesBuilder implements cog.Builder<prometheus.Scopes> {
    protected readonly internal: prometheus.Scopes;

    constructor() {
        this.internal = prometheus.defaultScopes();
    }

    /**
     * Builds the object.
     */
    build(): prometheus.Scopes {
        return this.internal;
    }

    defaultPath(defaultPath: string[]): this {
        this.internal.defaultPath = defaultPath;
        return this;
    }

    filters(filters: cog.Builder<prometheus.ScopesFilters>[]): this {
        const filtersResources = filters.map(builder1 => builder1.build());
        this.internal.filters = filtersResources;
        return this;
    }

    title(title: string): this {
        this.internal.title = title;
        return this;
    }
}

