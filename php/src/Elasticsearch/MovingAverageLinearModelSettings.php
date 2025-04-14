<?php

namespace Grafana\Foundation\Elasticsearch;

class MovingAverageLinearModelSettings implements \JsonSerializable
{
    public \Grafana\Foundation\Elasticsearch\MovingAverageModel $model;

    public string $window;

    public string $predict;

    /**
     * @param string|null $window
     * @param string|null $predict
     */
    public function __construct(?string $window = null, ?string $predict = null)
    {
        $this->model = \Grafana\Foundation\Elasticsearch\MovingAverageModel::simple();
        $this->window = $window ?: "";
        $this->predict = $predict ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{model?: "linear", window?: string, predict?: string} $inputData */
        $data = $inputData;
        return new self(
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
