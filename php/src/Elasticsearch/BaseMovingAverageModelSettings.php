<?php

namespace Grafana\Foundation\Elasticsearch;

class BaseMovingAverageModelSettings implements \JsonSerializable
{
    public \Grafana\Foundation\Elasticsearch\MovingAverageModel $model;

    public string $window;

    public string $predict;

    /**
     * @param \Grafana\Foundation\Elasticsearch\MovingAverageModel|null $model
     * @param string|null $window
     * @param string|null $predict
     */
    public function __construct(?\Grafana\Foundation\Elasticsearch\MovingAverageModel $model = null, ?string $window = null, ?string $predict = null)
    {
        $this->model = $model ?: \Grafana\Foundation\Elasticsearch\MovingAverageModel::Simple();
        $this->window = $window ?: "";
        $this->predict = $predict ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{model?: string, window?: string, predict?: string} $inputData */
        $data = $inputData;
        return new self(
            model: isset($data["model"]) ? (function($input) { return \Grafana\Foundation\Elasticsearch\MovingAverageModel::fromValue($input); })($data["model"]) : null,
            window: $data["window"] ?? null,
            predict: $data["predict"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "model" => $this->model,
            "window" => $this->window,
            "predict" => $this->predict,
        ];
        return $data;
    }
}
