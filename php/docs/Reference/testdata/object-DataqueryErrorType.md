---
title: <span class="badge object-type-enum"></span> DataqueryErrorType
---
# <span class="badge object-type-enum"></span> DataqueryErrorType

## Definition

```php
final class DataqueryErrorType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, DataqueryErrorType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function serverPanic(): self
    {
        if (!isset(self::$instances["ServerPanic"])) {
            self::$instances["ServerPanic"] = new self("server_panic");
        }

        return self::$instances["ServerPanic"];
    }

    public static function frontendException(): self
    {
        if (!isset(self::$instances["FrontendException"])) {
            self::$instances["FrontendException"] = new self("frontend_exception");
        }

        return self::$instances["FrontendException"];
    }

    public static function frontendObservable(): self
    {
        if (!isset(self::$instances["FrontendObservable"])) {
            self::$instances["FrontendObservable"] = new self("frontend_observable");
        }

        return self::$instances["FrontendObservable"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "server_panic") {
            return self::serverPanic();
        }

        if ($value === "frontend_exception") {
            return self::frontendException();
        }

        if ($value === "frontend_observable") {
            return self::frontendObservable();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum DataqueryErrorType");
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

```
