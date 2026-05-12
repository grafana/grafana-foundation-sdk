<?php

namespace Grafana\Foundation\Dashboardv2;

final class MatcherScope implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, MatcherScope>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function series(): self
    {
        if (!isset(self::$instances["series"])) {
            self::$instances["series"] = new self("series");
        }

        return self::$instances["series"];
    }

    public static function nested(): self
    {
        if (!isset(self::$instances["nested"])) {
            self::$instances["nested"] = new self("nested");
        }

        return self::$instances["nested"];
    }

    public static function annotation(): self
    {
        if (!isset(self::$instances["annotation"])) {
            self::$instances["annotation"] = new self("annotation");
        }

        return self::$instances["annotation"];
    }

    public static function exemplar(): self
    {
        if (!isset(self::$instances["exemplar"])) {
            self::$instances["exemplar"] = new self("exemplar");
        }

        return self::$instances["exemplar"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "series") {
            return self::series();
        }

        if ($value === "nested") {
            return self::nested();
        }

        if ($value === "annotation") {
            return self::annotation();
        }

        if ($value === "exemplar") {
            return self::exemplar();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum MatcherScope");
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

