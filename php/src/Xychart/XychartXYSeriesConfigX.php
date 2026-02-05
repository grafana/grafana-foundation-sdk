<?php

namespace Grafana\Foundation\Xychart;

class XychartXYSeriesConfigX implements \JsonSerializable
{
    public \Grafana\Foundation\Xychart\MatcherConfig $matcher;

    /**
     * @param \Grafana\Foundation\Xychart\MatcherConfig|null $matcher
     */
    public function __construct(?\Grafana\Foundation\Xychart\MatcherConfig $matcher = null)
    {
        $this->matcher = $matcher ?: new \Grafana\Foundation\Xychart\MatcherConfig();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{matcher?: mixed} $inputData */
        $data = $inputData;
        return new self(
            matcher: isset($data["matcher"]) ? (function($input) {
    	/** @var array{id?: string, options?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Xychart\MatcherConfig::fromArray($val);
    })($data["matcher"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->matcher = $this->matcher;
        return $data;
    }
}
