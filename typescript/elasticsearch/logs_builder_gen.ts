// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class LogsBuilder implements cog.OptionsBuilder<elasticsearch.Logs> {
    private readonly internal: elasticsearch.Logs;

    constructor() {
        this.internal = elasticsearch.defaultLogs();
        this.internal.type = "logs";
    }

    build(): elasticsearch.Logs {
        return this.internal;
    }

    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    settings(settings: {
	limit?: string;
}): this {
        this.internal.settings = settings;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }
}
