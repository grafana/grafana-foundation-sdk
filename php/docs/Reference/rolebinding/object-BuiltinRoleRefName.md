---
title: <span class="badge object-type-enum"></span> BuiltinRoleRefName
---
# <span class="badge object-type-enum"></span> BuiltinRoleRefName

## Definition

```php
final class BuiltinRoleRefName implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, BuiltinRoleRefName>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function viewer(): self
    {
        if (!isset(self::$instances["Viewer"])) {
            self::$instances["Viewer"] = new self("viewer");
        }

        return self::$instances["Viewer"];
    }

    public static function editor(): self
    {
        if (!isset(self::$instances["Editor"])) {
            self::$instances["Editor"] = new self("editor");
        }

        return self::$instances["Editor"];
    }

    public static function admin(): self
    {
        if (!isset(self::$instances["Admin"])) {
            self::$instances["Admin"] = new self("admin");
        }

        return self::$instances["Admin"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "viewer") {
            return self::viewer();
        }

        if ($value === "editor") {
            return self::editor();
        }

        if ($value === "admin") {
            return self::admin();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum BuiltinRoleRefName");
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
