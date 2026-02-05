<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class VizConfigKind implements \JsonSerializable
{
    public string $kind;

    /**
     * The group is the plugin ID
     */
    public string $group;

    public string $version;

    public \Grafana\Foundation\Dashboardv2beta1\VizConfigSpec $spec;

    /**
     * @param string|null $group
     * @param string|null $version
     * @param \Grafana\Foundation\Dashboardv2beta1\VizConfigSpec|null $spec
     */
    public function __construct(?string $group = null, ?string $version = null, ?\Grafana\Foundation\Dashboardv2beta1\VizConfigSpec $spec = null)
    {
        $this->kind = "VizConfig";
    
        $this->group = $group ?: "";
        $this->version = $version ?: "";
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\VizConfigSpec();
    }

    
    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{kind?: string, group?: string, version?: string, spec?: mixed} $inputData */
        $data = $inputData;
        return new self(
            group: $data["group"] ?? null,version: $data["version"] ?? null,spec: isset($data['spec']) ? (function($input) {
        /** @var array{options?: mixed, fieldConfig?: mixed} $spec */
        $spec = $input['spec'];
        /** @var string $group */
        $group = $input['group'] ?? '';
        return new VizConfigSpec(
            options: isset($spec["options"]) ? (function(string $panelType, $panel) {
        /** @var array<string, mixed> $options */
        $options = $panel["options"];
    
        if (!\Grafana\Foundation\Cog\Runtime::get()->panelcfgVariantExists($panelType)) {
            return $options;
        }
    
        $config = \Grafana\Foundation\Cog\Runtime::get()->panelcfgVariantConfig($panelType);
        if ($config->optionsFromArray === null) {
            return $options;
        }
    
        return ($config->optionsFromArray)($options);
    })($group, $spec) : null,
            fieldConfig: isset($spec["fieldConfig"]) ? (function(string $panelType, $panel) {
        /** @var array<string, mixed> $fieldConfigData */
        $fieldConfigData = $panel["fieldConfig"];
        $fieldConfigSource = FieldConfigSource::fromArray($fieldConfigData);
    
        if (!\Grafana\Foundation\Cog\Runtime::get()->panelcfgVariantExists($panelType)) {
            return $fieldConfigSource;
        }
    
        $config = \Grafana\Foundation\Cog\Runtime::get()->panelcfgVariantConfig($panelType);
        if ($config->fieldConfigFromArray === null) {
            return $fieldConfigSource;
        }
    
        if (!isset($fieldConfigData["defaults"])) {
            return $fieldConfigSource;
        }
    
        /** @var array{custom?: array<string, mixed>} $defaults */
        $defaults = $fieldConfigData["defaults"];
        if (!isset($defaults["custom"])) {
            return $fieldConfigSource;
        }
    
        $fieldConfigSource->defaults->custom = ($config->fieldConfigFromArray)($defaults["custom"]);
    
        return $fieldConfigSource;
    })($group, $spec) : null,
       );
    })($data) : null,
       );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->kind = $this->kind;
        $data->group = $this->group;
        $data->version = $this->version;
        $data->spec = $this->spec;
        return $data;
    }
}
