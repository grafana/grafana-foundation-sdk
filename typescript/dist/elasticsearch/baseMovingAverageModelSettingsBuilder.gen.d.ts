import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class BaseMovingAverageModelSettingsBuilder implements cog.Builder<elasticsearch.BaseMovingAverageModelSettings> {
    protected readonly internal: elasticsearch.BaseMovingAverageModelSettings;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.BaseMovingAverageModelSettings;
    model(model: elasticsearch.MovingAverageModel): this;
    window(window: string): this;
    predict(predict: string): this;
}
