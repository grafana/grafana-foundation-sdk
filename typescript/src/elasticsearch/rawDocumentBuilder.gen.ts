// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class RawDocumentBuilder implements cog.Builder<elasticsearch.RawDocument> {
    protected readonly internal: elasticsearch.RawDocument;

    constructor() {
        this.internal = elasticsearch.defaultRawDocument();
        this.internal.type = "raw_document";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.RawDocument {
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
