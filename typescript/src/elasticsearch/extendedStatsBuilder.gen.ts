// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class ExtendedStatsBuilder implements cog.Builder<elasticsearch.ExtendedStats> {
    protected readonly internal: elasticsearch.ExtendedStats;

    constructor() {
        this.internal = elasticsearch.defaultExtendedStats();
        this.internal.type = "extended_stats";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.ExtendedStats {
        return this.internal;
    }

    settings(settings: {
	script?: elasticsearch.InlineScript;
	missing?: string;
	sigma?: string;
}): this {
        this.internal.settings = settings;
        return this;
    }

    field(field: string): this {
        this.internal.field = field;
        return this;
    }

    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    meta(meta: any): this {
        this.internal.meta = meta;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }
}
