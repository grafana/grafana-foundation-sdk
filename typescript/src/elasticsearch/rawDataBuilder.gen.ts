// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class RawDataBuilder implements cog.Builder<elasticsearch.RawData> {
    protected readonly internal: elasticsearch.RawData;

    constructor() {
        this.internal = elasticsearch.defaultRawData();
        this.internal.type = "raw_data";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.RawData {
        return this.internal;
    }

    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    settings(settings: {
	size?: string;
}): this {
        this.internal.settings = settings;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }
}
