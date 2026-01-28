// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as news from '../news';

export class OptionsBuilder implements cog.Builder<news.Options> {
    protected readonly internal: news.Options;

    constructor() {
        this.internal = news.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): news.Options {
        return this.internal;
    }

    // empty/missing will default to grafana blog
    feedUrl(feedUrl: string): this {
        this.internal.feedUrl = feedUrl;
        return this;
    }

    showImage(showImage: boolean): this {
        this.internal.showImage = showImage;
        return this;
    }
}

