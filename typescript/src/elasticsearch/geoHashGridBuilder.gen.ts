// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class GeoHashGridBuilder implements cog.Builder<elasticsearch.GeoHashGrid> {
    protected readonly internal: elasticsearch.GeoHashGrid;

    constructor() {
        this.internal = elasticsearch.defaultGeoHashGrid();
        this.internal.type = "geohash_grid";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.GeoHashGrid {
        return this.internal;
    }

    field(field: string): this {
        this.internal.field = field;
        return this;
    }

    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    settings(settings: {
	precision?: string;
}): this {
        this.internal.settings = settings;
        return this;
    }
}
