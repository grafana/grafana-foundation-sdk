<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchMovingAverageHoltWintersModelSettingsSettings implements \JsonSerializable
{
    public ?string $alpha;

    public ?string $beta;

    public ?string $gamma;

    public ?string $period;

    public ?bool $pad;

    /**
     * @param string|null $alpha
     * @param string|null $beta
     * @param string|null $gamma
     * @param string|null $period
     * @param bool|null $pad
     */
    public function __construct(?string $alpha = null, ?string $beta = null, ?string $gamma = null, ?string $period = null, ?bool $pad = null)
    {
        $this->alpha = $alpha;
        $this->beta = $beta;
        $this->gamma = $gamma;
        $this->period = $period;
        $this->pad = $pad;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{alpha?: string, beta?: string, gamma?: string, period?: string, pad?: bool} $inputData */
        $data = $inputData;
        return new self(
            alpha: $data["alpha"] ?? null,
            beta: $data["beta"] ?? null,
            gamma: $data["gamma"] ?? null,
            period: $data["period"] ?? null,
            pad: $data["pad"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->alpha)) {
            $data["alpha"] = $this->alpha;
        }
        if (isset($this->beta)) {
            $data["beta"] = $this->beta;
        }
        if (isset($this->gamma)) {
            $data["gamma"] = $this->gamma;
        }
        if (isset($this->period)) {
            $data["period"] = $this->period;
        }
        if (isset($this->pad)) {
            $data["pad"] = $this->pad;
        }
        return $data;
    }
}
