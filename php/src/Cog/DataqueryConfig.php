<?php

namespace Grafana\Foundation\Cog;

final class DataqueryConfig
{
    public readonly string $identifier;

    /**
     * @var callable(array<string, mixed>): Dataquery
     */
    public $fromArray;

    /**
     * @param callable(array<string, mixed>): Dataquery $fromArray
     */
    public function __construct(string $identifier, callable $fromArray)
    {
        $this->identifier = $identifier;
        $this->fromArray = $fromArray;
    }
}
