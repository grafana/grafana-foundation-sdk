// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv0alpha1 from '../alertingv0alpha1';

/**
 * #SearchResultsMetadata carries the paging token and total authorised match
 * count.
 */
export class SearchResultsMetadataBuilder implements cog.Builder<alertingv0alpha1.SearchResultsMetadata> {
    protected readonly internal: alertingv0alpha1.SearchResultsMetadata;

    constructor() {
        this.internal = alertingv0alpha1.defaultSearchResultsMetadata();
    }

    /**
     * Builds the object.
     */
    build(): alertingv0alpha1.SearchResultsMetadata {
        return this.internal;
    }

    continueVal(continueVal: string): this {
        this.internal.continue = continueVal;
        return this;
    }

    totalHits(totalHits: number): this {
        this.internal.totalHits = totalHits;
        return this;
    }

    // Always read totalHits together with totalHitsRelation.
    totalHitsRelation(totalHitsRelation: alertingv0alpha1.TotalHitsRelation): this {
        this.internal.totalHitsRelation = totalHitsRelation;
        return this;
    }
}

