// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv0alpha1 from '../alertingv0alpha1';

/**
 * #SearchResultHit is a single match: its identity, an optional relevance
 * score (present only when the query included free text), and the requested
 * fields.
 */
export class SearchResultHitBuilder implements cog.Builder<alertingv0alpha1.SearchResultHit> {
    protected readonly internal: alertingv0alpha1.SearchResultHit;

    constructor() {
        this.internal = alertingv0alpha1.defaultSearchResultHit();
    }

    /**
     * Builds the object.
     */
    build(): alertingv0alpha1.SearchResultHit {
        return this.internal;
    }

    resource(resource: cog.Builder<alertingv0alpha1.SearchResultResource>): this {
        const resourceResource = resource.build();
        this.internal.resource = resourceResource;
        return this;
    }

    score(score: number): this {
        this.internal.score = score;
        return this;
    }

    fields(fields: cog.Builder<alertingv0alpha1.RuleSearchHitFields>): this {
        const fieldsResource = fields.build();
        this.internal.fields = fieldsResource;
        return this;
    }
}

