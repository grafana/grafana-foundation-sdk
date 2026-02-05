import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class LogsBuilder implements cog.Builder<elasticsearch.Logs> {
    protected readonly internal: elasticsearch.Logs;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.Logs;
    id(id: string): this;
    settings(settings: {
        limit?: string;
    }): this;
    hide(hide: boolean): this;
}
