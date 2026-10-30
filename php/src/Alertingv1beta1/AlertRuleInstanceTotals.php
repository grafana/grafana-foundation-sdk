<?php

namespace Grafana\Foundation\Alertingv1beta1;

/**
 * Count of alert instances per state. error also counts instances whose evaluation
 * errored but were mapped to another state via execErrState, so it can overlap.
 */
class AlertRuleInstanceTotals implements \JsonSerializable
{
    public ?int $healthy;

    public ?int $firing;

    public ?int $pending;

    public ?int $recovering;

    public ?int $nodata;

    public ?int $error;

    /**
     * @param int|null $healthy
     * @param int|null $firing
     * @param int|null $pending
     * @param int|null $recovering
     * @param int|null $nodata
     * @param int|null $error
     */
    public function __construct(?int $healthy = null, ?int $firing = null, ?int $pending = null, ?int $recovering = null, ?int $nodata = null, ?int $error = null)
    {
        $this->healthy = $healthy;
        $this->firing = $firing;
        $this->pending = $pending;
        $this->recovering = $recovering;
        $this->nodata = $nodata;
        $this->error = $error;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{healthy?: int, firing?: int, pending?: int, recovering?: int, nodata?: int, error?: int} $inputData */
        $data = $inputData;
        return new self(
            healthy: $data["healthy"] ?? null,
            firing: $data["firing"] ?? null,
            pending: $data["pending"] ?? null,
            recovering: $data["recovering"] ?? null,
            nodata: $data["nodata"] ?? null,
            error: $data["error"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->healthy)) {
            $data->healthy = $this->healthy;
        }
        if (isset($this->firing)) {
            $data->firing = $this->firing;
        }
        if (isset($this->pending)) {
            $data->pending = $this->pending;
        }
        if (isset($this->recovering)) {
            $data->recovering = $this->recovering;
        }
        if (isset($this->nodata)) {
            $data->nodata = $this->nodata;
        }
        if (isset($this->error)) {
            $data->error = $this->error;
        }
        return $data;
    }
}
