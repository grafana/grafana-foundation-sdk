<?php

namespace Grafana\Foundation\Elasticsearch;

class MovingAverageEWMAModelSettings implements \JsonSerializable
{
    public \Grafana\Foundation\Elasticsearch\MovingAverageModel $model;

    public ?\Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageEWMAModelSettingsSettings $settings;

    public string $window;

    public bool $minimize;

    public string $predict;

    /**
     * @param \Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageEWMAModelSettingsSettings|null $settings
     * @param string|null $window
     * @param bool|null $minimize
     * @param string|null $predict
     */
    public function __construct(?\Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageEWMAModelSettingsSettings $settings = null, ?string $window = null, ?bool $minimize = null, ?string $predict = null)
    {
        $this->model = \Grafana\Foundation\Elasticsearch\MovingAverageModel::simple();
        $this->settings = $settings;
        $this->window = $window ?: "";
        $this->minimize = $minimize ?: false;
        $this->predict = $predict ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{model?: "ewma", settings?: mixed, window?: string, minimize?: bool, predict?: string} $inputData */
        $data = $inputData;
        return new self(
            settings: isset($data["settings"]) ? (function($input) {
    	/** @var array{alpha?: string} */
    $val = $input;
    	return \Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageEWMAModelSettingsSettings::fromArray($val);
    })($data["settings"]) : null,
            window: $data["window"] ?? null,
            minimize: $data["minimize"] ?? null,
            predict: $data["predict"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->model = $this->model;
        $data->window = $this->window;
        $data->minimize = $this->minimize;
        $data->predict = $this->predict;
        if (isset($this->settings)) {
            $data->settings = $this->settings;
        }
        return $data;
    }
}
