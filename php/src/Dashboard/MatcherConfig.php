<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.
 * It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.
 */
class MatcherConfig implements \JsonSerializable
{
    /**
     * The matcher id. This is used to find the matcher implementation from registry.
     */
    public string $id;

    /**
     * The matcher options. This is specific to the matcher implementation.
     * @var mixed|null
     */
    public $options;

    /**
     * @param string|null $id
     * @param mixed|null $options
     */
    public function __construct(?string $id = null,  $options = null)
    {
        $this->id = $id ?: "";
        $this->options = $options;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{id?: string, options?: mixed} $inputData */
        $data = $inputData;
        return new self(
            id: $data["id"] ?? null,
            options: $data["options"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "id" => $this->id,
        ];
        if (isset($this->options)) {
            $data["options"] = $this->options;
        }
        return $data;
    }
}
