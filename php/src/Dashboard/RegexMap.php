<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Maps regular expressions to replacement text and a color.
 * For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
 */
class RegexMap implements \JsonSerializable
{
    public string $type;

    /**
     * Regular expression to match against and the result to apply when the value matches the regex
     */
    public \Grafana\Foundation\Dashboard\DashboardRegexMapOptions $options;

    /**
     * @param \Grafana\Foundation\Dashboard\DashboardRegexMapOptions|null $options
     */
    public function __construct(?\Grafana\Foundation\Dashboard\DashboardRegexMapOptions $options = null)
    {
        $this->type = "regex";
    
        $this->options = $options ?: new \Grafana\Foundation\Dashboard\DashboardRegexMapOptions();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, options?: mixed} $inputData */
        $data = $inputData;
        return new self(
            options: isset($data["options"]) ? (function($input) {
    	/** @var array{pattern?: string, result?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DashboardRegexMapOptions::fromArray($val);
    })($data["options"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "type" => $this->type,
            "options" => $this->options,
        ];
        return $data;
    }
}
