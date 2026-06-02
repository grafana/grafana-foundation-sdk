<?php

namespace Grafana\Foundation\Loki;

final class LokiQueryDirection implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, LokiQueryDirection>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function forward(): self
    {
        if (!isset(self::$instances["Forward"])) {
            self::$instances["Forward"] = new self("forward");
        }

        return self::$instances["Forward"];
    }

    public static function backward(): self
    {
        if (!isset(self::$instances["Backward"])) {
            self::$instances["Backward"] = new self("backward");
        }

        return self::$instances["Backward"];
    }

    public static function scan(): self
    {
        if (!isset(self::$instances["Scan"])) {
            self::$instances["Scan"] = new self("scan");
        }

        return self::$instances["Scan"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "forward") {
            return self::forward();
        }

        if ($value === "backward") {
            return self::backward();
        }

        if ($value === "scan") {
            return self::scan();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum LokiQueryDirection");
    }

    public function jsonSerialize(): string
    {
        return $this->value;
    }

    public function __toString(): string
    {
        return $this->value;
    }
}

