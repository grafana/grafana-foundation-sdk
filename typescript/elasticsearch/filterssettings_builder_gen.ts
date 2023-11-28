// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class FiltersSettingsBuilder implements cog.Builder<elasticsearch.FiltersSettings> {
    private readonly internal: elasticsearch.FiltersSettings;

    constructor() {
        this.internal = elasticsearch.defaultFiltersSettings();
    }

    build(): elasticsearch.FiltersSettings {
        return this.internal;
    }

    filters(filters: cog.Builder<elasticsearch.Filter>[]): this {
        const filtersResources = filters.map(builder => builder.build());
        this.internal.filters = filtersResources;
        return this;
    }
}
