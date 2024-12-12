<?php

namespace {{ .Data.NamespaceRoot }}\Cog;

final class DataqueryConfig
{
    public readonly string $identifier;

    /**
     * @var callable(array<string, mixed>): Dataquery
     */
    public $fromArray;

    /**
     * @var (callable(Dataquery): string)|null
     */
    public $convert;

    /**
     * @param callable(array<string, mixed>): Dataquery $fromArray
     * @param (callable(Dataquery): string)|null $convert
     */
    public function __construct(string $identifier, callable $fromArray, ?callable $convert = null)
    {
        $this->identifier = $identifier;
        $this->fromArray = $fromArray;
        $this->convert = $convert;
    }
}
