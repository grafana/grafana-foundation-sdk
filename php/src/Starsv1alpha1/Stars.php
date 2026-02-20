<?php

namespace Grafana\Foundation\Starsv1alpha1;

class Stars implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Starsv1alpha1\Resource>
     */
    public array $resource;

    /**
     * @param array<\Grafana\Foundation\Starsv1alpha1\Resource>|null $resource
     */
    public function __construct(?array $resource = null)
    {
        $this->resource = $resource ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{resource?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            resource: array_filter(array_map((function($input) {
    	/** @var array{group?: string, kind?: string, names?: array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Starsv1alpha1\Resource::fromArray($val);
    }), $data["resource"] ?? [])),
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->resource = $this->resource;
        return $data;
    }
}
