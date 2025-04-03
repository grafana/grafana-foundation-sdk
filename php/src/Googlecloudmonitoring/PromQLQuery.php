<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

/**
 * PromQL sub-query properties.
 */
class PromQLQuery implements \JsonSerializable
{
    /**
     * GCP project to execute the query against.
     */
    public string $projectName;

    /**
     * PromQL expression/query to be executed.
     */
    public string $expr;

    /**
     * PromQL min step
     */
    public string $step;

    /**
     * @param string|null $projectName
     * @param string|null $expr
     * @param string|null $step
     */
    public function __construct(?string $projectName = null, ?string $expr = null, ?string $step = null)
    {
        $this->projectName = $projectName ?: "";
        $this->expr = $expr ?: "";
        $this->step = $step ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{projectName?: string, expr?: string, step?: string} $inputData */
        $data = $inputData;
        return new self(
            projectName: $data["projectName"] ?? null,
            expr: $data["expr"] ?? null,
            step: $data["step"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "projectName" => $this->projectName,
            "expr" => $this->expr,
            "step" => $this->step,
        ];
        return $data;
    }
}
