// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class FiltersBuilder implements cog.Builder<elasticsearch.Filters> {
    protected readonly internal: elasticsearch.Filters;

    constructor() {
        this.internal = elasticsearch.defaultFilters();
        this.internal.type = "filters";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.Filters {
        return this.internal;
    }

    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    settings(settings: {
	filters?: elasticsearch.Filter[];
}): this {
        this.internal.settings = settings;
        return this;
    }
}
