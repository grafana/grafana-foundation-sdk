// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class TermsBuilder implements cog.Builder<elasticsearch.Terms> {
    protected readonly internal: elasticsearch.Terms;

    constructor() {
        this.internal = elasticsearch.defaultTerms();
        this.internal.type = "terms";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.Terms {
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
	order?: elasticsearch.TermsOrder;
	size?: string;
	min_doc_count?: string;
	orderBy?: string;
	missing?: string;
}): this {
        this.internal.settings = settings;
        return this;
    }
}
