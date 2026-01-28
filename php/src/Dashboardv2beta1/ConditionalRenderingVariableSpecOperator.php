<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class ConditionalRenderingVariableSpecOperator implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ConditionalRenderingVariableSpecOperator>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function equals(): self
    {
        if (!isset(self::$instances["Equals"])) {
            self::$instances["Equals"] = new self("equals");
        }

        return self::$instances["Equals"];
    }

    public static function notEquals(): self
    {
        if (!isset(self::$instances["NotEquals"])) {
            self::$instances["NotEquals"] = new self("notEquals");
        }

        return self::$instances["NotEquals"];
    }

    public static function matches(): self
    {
        if (!isset(self::$instances["Matches"])) {
            self::$instances["Matches"] = new self("matches");
        }

        return self::$instances["Matches"];
    }

    public static function notMatches(): self
    {
        if (!isset(self::$instances["NotMatches"])) {
            self::$instances["NotMatches"] = new self("notMatches");
        }

        return self::$instances["NotMatches"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "equals") {
            return self::equals();
        }

        if ($value === "notEquals") {
            return self::notEquals();
        }

        if ($value === "matches") {
            return self::matches();
        }

        if ($value === "notMatches") {
            return self::notMatches();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ConditionalRenderingVariableSpecOperator");
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

