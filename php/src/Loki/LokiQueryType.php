<?php

namespace Grafana\Foundation\Loki;

final class LokiQueryType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, LokiQueryType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function range(): self
    {
        if (!isset(self::$instances["Range"])) {
            self::$instances["Range"] = new self("range");
        }

        return self::$instances["Range"];
    }

    public static function instant(): self
    {
        if (!isset(self::$instances["Instant"])) {
            self::$instances["Instant"] = new self("instant");
        }

        return self::$instances["Instant"];
    }

    public static function stream(): self
    {
        if (!isset(self::$instances["Stream"])) {
            self::$instances["Stream"] = new self("stream");
        }

        return self::$instances["Stream"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "range") {
            return self::range();
        }

        if ($value === "instant") {
            return self::instant();
        }

        if ($value === "stream") {
            return self::stream();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum LokiQueryType");
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

