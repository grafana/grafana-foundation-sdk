<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class VizConfigKindConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\VizConfigKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\VizConfigKindBuilder())',
        ];
            if ($input->group !== "") {
    
        
    $buffer = 'group(';
        $arg0 =\var_export($input->group, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->version !== "") {
    
        
    $buffer = 'version(';
        $arg0 =\var_export($input->version, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'spec(';
        $arg0 ='(new \Grafana\Foundation\Dashboardv2beta1\VizConfigSpec('.(($input->spec->options !== null) ? 'options: '.\var_export($input->spec->options, true).', ' : '').'fieldConfig: '.'(new \Grafana\Foundation\Dashboardv2beta1\FieldConfigSource(defaults: '.'(new \Grafana\Foundation\Dashboardv2beta1\FieldConfig('.(($input->spec->fieldConfig->defaults->displayName !== null) ? 'displayName: '.\var_export($input->spec->fieldConfig->defaults->displayName, true).', ' : '').''.(($input->spec->fieldConfig->defaults->displayNameFromDS !== null) ? 'displayNameFromDS: '.\var_export($input->spec->fieldConfig->defaults->displayNameFromDS, true).', ' : '').''.(($input->spec->fieldConfig->defaults->description !== null) ? 'description: '.\var_export($input->spec->fieldConfig->defaults->description, true).', ' : '').''.(($input->spec->fieldConfig->defaults->path !== null) ? 'path: '.\var_export($input->spec->fieldConfig->defaults->path, true).', ' : '').''.(($input->spec->fieldConfig->defaults->writeable !== null) ? 'writeable: '.\var_export($input->spec->fieldConfig->defaults->writeable, true).', ' : '').''.(($input->spec->fieldConfig->defaults->filterable !== null) ? 'filterable: '.\var_export($input->spec->fieldConfig->defaults->filterable, true).', ' : '').''.(($input->spec->fieldConfig->defaults->unit !== null) ? 'unit: '.\var_export($input->spec->fieldConfig->defaults->unit, true).', ' : '').''.(($input->spec->fieldConfig->defaults->decimals !== null) ? 'decimals: '.\var_export($input->spec->fieldConfig->defaults->decimals, true).', ' : '').''.(($input->spec->fieldConfig->defaults->min !== null) ? 'min: '.\var_export($input->spec->fieldConfig->defaults->min, true).', ' : '').''.(($input->spec->fieldConfig->defaults->max !== null) ? 'max: '.\var_export($input->spec->fieldConfig->defaults->max, true).', ' : '').''.(($input->spec->fieldConfig->defaults->mappings !== null) ? 'mappings: '.\var_export($input->spec->fieldConfig->defaults->mappings, true).', ' : '').''.(($input->spec->fieldConfig->defaults->thresholds !== null) ? 'thresholds: '.'(new \Grafana\Foundation\Dashboardv2beta1\ThresholdsConfig(mode: '.'\Grafana\Foundation\Dashboardv2beta1\ThresholdsMode::fromValue("'.$input->spec->fieldConfig->defaults->thresholds->mode.'")'.',steps: '.\var_export($input->spec->fieldConfig->defaults->thresholds->steps, true).',))'.', ' : '').''.(($input->spec->fieldConfig->defaults->color !== null) ? 'color: '.'(new \Grafana\Foundation\Dashboardv2beta1\FieldColor(mode: '.'\Grafana\Foundation\Dashboardv2beta1\FieldColorModeId::fromValue("'.$input->spec->fieldConfig->defaults->color->mode.'")'.','.(($input->spec->fieldConfig->defaults->color->fixedColor !== null) ? 'fixedColor: '.\var_export($input->spec->fieldConfig->defaults->color->fixedColor, true).', ' : '').''.(($input->spec->fieldConfig->defaults->color->seriesBy !== null) ? 'seriesBy: '.'\Grafana\Foundation\Dashboardv2beta1\FieldColorSeriesByMode::fromValue("'.$input->spec->fieldConfig->defaults->color->seriesBy.'")'.', ' : '').'))'.', ' : '').''.(($input->spec->fieldConfig->defaults->links !== null) ? 'links: '.\var_export($input->spec->fieldConfig->defaults->links, true).', ' : '').''.(($input->spec->fieldConfig->defaults->actions !== null) ? 'actions: '.\var_export($input->spec->fieldConfig->defaults->actions, true).', ' : '').''.(($input->spec->fieldConfig->defaults->noValue !== null) ? 'noValue: '.\var_export($input->spec->fieldConfig->defaults->noValue, true).', ' : '').''.(($input->spec->fieldConfig->defaults->custom !== null) ? 'custom: '.\var_export($input->spec->fieldConfig->defaults->custom, true).', ' : '').'))'.',overrides: '.\var_export($input->spec->fieldConfig->overrides, true).',))'.',))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec->options !== null) {
    
        
    $buffer = 'options(';
        $arg0 =\var_export($input->spec->options, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'fieldConfig(';
        $arg0 ='(new \Grafana\Foundation\Dashboardv2beta1\FieldConfigSource(defaults: '.'(new \Grafana\Foundation\Dashboardv2beta1\FieldConfig('.(($input->spec->fieldConfig->defaults->displayName !== null) ? 'displayName: '.\var_export($input->spec->fieldConfig->defaults->displayName, true).', ' : '').''.(($input->spec->fieldConfig->defaults->displayNameFromDS !== null) ? 'displayNameFromDS: '.\var_export($input->spec->fieldConfig->defaults->displayNameFromDS, true).', ' : '').''.(($input->spec->fieldConfig->defaults->description !== null) ? 'description: '.\var_export($input->spec->fieldConfig->defaults->description, true).', ' : '').''.(($input->spec->fieldConfig->defaults->path !== null) ? 'path: '.\var_export($input->spec->fieldConfig->defaults->path, true).', ' : '').''.(($input->spec->fieldConfig->defaults->writeable !== null) ? 'writeable: '.\var_export($input->spec->fieldConfig->defaults->writeable, true).', ' : '').''.(($input->spec->fieldConfig->defaults->filterable !== null) ? 'filterable: '.\var_export($input->spec->fieldConfig->defaults->filterable, true).', ' : '').''.(($input->spec->fieldConfig->defaults->unit !== null) ? 'unit: '.\var_export($input->spec->fieldConfig->defaults->unit, true).', ' : '').''.(($input->spec->fieldConfig->defaults->decimals !== null) ? 'decimals: '.\var_export($input->spec->fieldConfig->defaults->decimals, true).', ' : '').''.(($input->spec->fieldConfig->defaults->min !== null) ? 'min: '.\var_export($input->spec->fieldConfig->defaults->min, true).', ' : '').''.(($input->spec->fieldConfig->defaults->max !== null) ? 'max: '.\var_export($input->spec->fieldConfig->defaults->max, true).', ' : '').''.(($input->spec->fieldConfig->defaults->mappings !== null) ? 'mappings: '.\var_export($input->spec->fieldConfig->defaults->mappings, true).', ' : '').''.(($input->spec->fieldConfig->defaults->thresholds !== null) ? 'thresholds: '.'(new \Grafana\Foundation\Dashboardv2beta1\ThresholdsConfig(mode: '.'\Grafana\Foundation\Dashboardv2beta1\ThresholdsMode::fromValue("'.$input->spec->fieldConfig->defaults->thresholds->mode.'")'.',steps: '.\var_export($input->spec->fieldConfig->defaults->thresholds->steps, true).',))'.', ' : '').''.(($input->spec->fieldConfig->defaults->color !== null) ? 'color: '.'(new \Grafana\Foundation\Dashboardv2beta1\FieldColor(mode: '.'\Grafana\Foundation\Dashboardv2beta1\FieldColorModeId::fromValue("'.$input->spec->fieldConfig->defaults->color->mode.'")'.','.(($input->spec->fieldConfig->defaults->color->fixedColor !== null) ? 'fixedColor: '.\var_export($input->spec->fieldConfig->defaults->color->fixedColor, true).', ' : '').''.(($input->spec->fieldConfig->defaults->color->seriesBy !== null) ? 'seriesBy: '.'\Grafana\Foundation\Dashboardv2beta1\FieldColorSeriesByMode::fromValue("'.$input->spec->fieldConfig->defaults->color->seriesBy.'")'.', ' : '').'))'.', ' : '').''.(($input->spec->fieldConfig->defaults->links !== null) ? 'links: '.\var_export($input->spec->fieldConfig->defaults->links, true).', ' : '').''.(($input->spec->fieldConfig->defaults->actions !== null) ? 'actions: '.\var_export($input->spec->fieldConfig->defaults->actions, true).', ' : '').''.(($input->spec->fieldConfig->defaults->noValue !== null) ? 'noValue: '.\var_export($input->spec->fieldConfig->defaults->noValue, true).', ' : '').''.(($input->spec->fieldConfig->defaults->custom !== null) ? 'custom: '.\var_export($input->spec->fieldConfig->defaults->custom, true).', ' : '').'))'.',overrides: '.\var_export($input->spec->fieldConfig->overrides, true).',))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec->fieldConfig->defaults->displayName !== null && $input->spec->fieldConfig->defaults->displayName !== "") {
    
        
    $buffer = 'displayName(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->displayName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->displayNameFromDS !== null && $input->spec->fieldConfig->defaults->displayNameFromDS !== "") {
    
        
    $buffer = 'displayNameFromDS(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->displayNameFromDS, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->description !== null && $input->spec->fieldConfig->defaults->description !== "") {
    
        
    $buffer = 'description(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->description, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->path !== null && $input->spec->fieldConfig->defaults->path !== "") {
    
        
    $buffer = 'path(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->path, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->writeable !== null) {
    
        
    $buffer = 'writeable(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->writeable, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->filterable !== null) {
    
        
    $buffer = 'filterable(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->filterable, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->unit !== null && $input->spec->fieldConfig->defaults->unit !== "") {
    
        
    $buffer = 'unit(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->unit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->decimals !== null) {
    
        
    $buffer = 'decimals(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->decimals, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->min !== null) {
    
        
    $buffer = 'min(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->min, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->max !== null) {
    
        
    $buffer = 'max(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->max, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->mappings !== null && count($input->spec->fieldConfig->defaults->mappings) >= 1) {
    
        
    $buffer = 'mappings(';
        $tmparg0 = [];
        foreach ($input->spec->fieldConfig->defaults->mappings as $arg1) {
        switch (true) {
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\ValueMap:
                $disjunctionarg1 ='(new \Grafana\Foundation\Dashboardv2beta1\ValueMap(type: '.\var_export($arg1->type, true).',options: '.\var_export($arg1->options, true).',))';
                $tmpmappingsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\RangeMap:
                $disjunctionarg1 ='(new \Grafana\Foundation\Dashboardv2beta1\RangeMap(type: '.\var_export($arg1->type, true).',options: '.'(new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RangeMapOptions('.(($arg1->options->from !== null) ? 'from: '.\var_export($arg1->options->from, true).', ' : '').''.(($arg1->options->to !== null) ? 'to: '.\var_export($arg1->options->to, true).', ' : '').'result: '.'(new \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult('.(($arg1->options->result->text !== null) ? 'text: '.\var_export($arg1->options->result->text, true).', ' : '').''.(($arg1->options->result->color !== null) ? 'color: '.\var_export($arg1->options->result->color, true).', ' : '').''.(($arg1->options->result->icon !== null) ? 'icon: '.\var_export($arg1->options->result->icon, true).', ' : '').''.(($arg1->options->result->index !== null) ? 'index: '.\var_export($arg1->options->result->index, true).', ' : '').'))'.',))'.',))';
                $tmpmappingsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\RegexMap:
                $disjunctionarg1 ='(new \Grafana\Foundation\Dashboardv2beta1\RegexMap(type: '.\var_export($arg1->type, true).',options: '.'(new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RegexMapOptions(pattern: '.\var_export($arg1->options->pattern, true).',result: '.'(new \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult('.(($arg1->options->result->text !== null) ? 'text: '.\var_export($arg1->options->result->text, true).', ' : '').''.(($arg1->options->result->color !== null) ? 'color: '.\var_export($arg1->options->result->color, true).', ' : '').''.(($arg1->options->result->icon !== null) ? 'icon: '.\var_export($arg1->options->result->icon, true).', ' : '').''.(($arg1->options->result->index !== null) ? 'index: '.\var_export($arg1->options->result->index, true).', ' : '').'))'.',))'.',))';
                $tmpmappingsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\SpecialValueMap:
                $disjunctionarg1 ='(new \Grafana\Foundation\Dashboardv2beta1\SpecialValueMap(type: '.\var_export($arg1->type, true).',options: '.'(new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1SpecialValueMapOptions(match: '.'\Grafana\Foundation\Dashboardv2beta1\SpecialValueMatch::fromValue("'.$arg1->options->match.'")'.',result: '.'(new \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult('.(($arg1->options->result->text !== null) ? 'text: '.\var_export($arg1->options->result->text, true).', ' : '').''.(($arg1->options->result->color !== null) ? 'color: '.\var_export($arg1->options->result->color, true).', ' : '').''.(($arg1->options->result->icon !== null) ? 'icon: '.\var_export($arg1->options->result->icon, true).', ' : '').''.(($arg1->options->result->index !== null) ? 'index: '.\var_export($arg1->options->result->index, true).', ' : '').'))'.',))'.',))';
                $tmpmappingsarg1 = $disjunctionarg1;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $tmparg0[] = $tmpmappingsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->thresholds !== null) {
    
        
    $buffer = 'thresholds(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\ThresholdsConfigConverter::convert($input->spec->fieldConfig->defaults->thresholds);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->color !== null) {
    
        
    $buffer = 'colorScheme(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\FieldColorConverter::convert($input->spec->fieldConfig->defaults->color);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->links !== null && count($input->spec->fieldConfig->defaults->links) >= 1) {
    
        
    $buffer = 'dataLinks(';
        $tmparg0 = [];
        foreach ($input->spec->fieldConfig->defaults->links as $arg1) {
        $tmplinksarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmplinksarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->actions !== null && count($input->spec->fieldConfig->defaults->actions) >= 1) {
    
        
    $buffer = 'actions(';
        $tmparg0 = [];
        foreach ($input->spec->fieldConfig->defaults->actions as $arg1) {
        $tmpactionsarg1 = \Grafana\Foundation\Dashboardv2beta1\ActionConverter::convert($arg1);
        $tmparg0[] = $tmpactionsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->noValue !== null && $input->spec->fieldConfig->defaults->noValue !== "") {
    
        
    $buffer = 'noValue(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->noValue, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null) {
    
        
    $buffer = 'custom(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->custom, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'defaults(';
        $arg0 ='(new \Grafana\Foundation\Dashboardv2beta1\FieldConfig('.(($input->spec->fieldConfig->defaults->displayName !== null) ? 'displayName: '.\var_export($input->spec->fieldConfig->defaults->displayName, true).', ' : '').''.(($input->spec->fieldConfig->defaults->displayNameFromDS !== null) ? 'displayNameFromDS: '.\var_export($input->spec->fieldConfig->defaults->displayNameFromDS, true).', ' : '').''.(($input->spec->fieldConfig->defaults->description !== null) ? 'description: '.\var_export($input->spec->fieldConfig->defaults->description, true).', ' : '').''.(($input->spec->fieldConfig->defaults->path !== null) ? 'path: '.\var_export($input->spec->fieldConfig->defaults->path, true).', ' : '').''.(($input->spec->fieldConfig->defaults->writeable !== null) ? 'writeable: '.\var_export($input->spec->fieldConfig->defaults->writeable, true).', ' : '').''.(($input->spec->fieldConfig->defaults->filterable !== null) ? 'filterable: '.\var_export($input->spec->fieldConfig->defaults->filterable, true).', ' : '').''.(($input->spec->fieldConfig->defaults->unit !== null) ? 'unit: '.\var_export($input->spec->fieldConfig->defaults->unit, true).', ' : '').''.(($input->spec->fieldConfig->defaults->decimals !== null) ? 'decimals: '.\var_export($input->spec->fieldConfig->defaults->decimals, true).', ' : '').''.(($input->spec->fieldConfig->defaults->min !== null) ? 'min: '.\var_export($input->spec->fieldConfig->defaults->min, true).', ' : '').''.(($input->spec->fieldConfig->defaults->max !== null) ? 'max: '.\var_export($input->spec->fieldConfig->defaults->max, true).', ' : '').''.(($input->spec->fieldConfig->defaults->mappings !== null) ? 'mappings: '.\var_export($input->spec->fieldConfig->defaults->mappings, true).', ' : '').''.(($input->spec->fieldConfig->defaults->thresholds !== null) ? 'thresholds: '.'(new \Grafana\Foundation\Dashboardv2beta1\ThresholdsConfig(mode: '.'\Grafana\Foundation\Dashboardv2beta1\ThresholdsMode::fromValue("'.$input->spec->fieldConfig->defaults->thresholds->mode.'")'.',steps: '.\var_export($input->spec->fieldConfig->defaults->thresholds->steps, true).',))'.', ' : '').''.(($input->spec->fieldConfig->defaults->color !== null) ? 'color: '.'(new \Grafana\Foundation\Dashboardv2beta1\FieldColor(mode: '.'\Grafana\Foundation\Dashboardv2beta1\FieldColorModeId::fromValue("'.$input->spec->fieldConfig->defaults->color->mode.'")'.','.(($input->spec->fieldConfig->defaults->color->fixedColor !== null) ? 'fixedColor: '.\var_export($input->spec->fieldConfig->defaults->color->fixedColor, true).', ' : '').''.(($input->spec->fieldConfig->defaults->color->seriesBy !== null) ? 'seriesBy: '.'\Grafana\Foundation\Dashboardv2beta1\FieldColorSeriesByMode::fromValue("'.$input->spec->fieldConfig->defaults->color->seriesBy.'")'.', ' : '').'))'.', ' : '').''.(($input->spec->fieldConfig->defaults->links !== null) ? 'links: '.\var_export($input->spec->fieldConfig->defaults->links, true).', ' : '').''.(($input->spec->fieldConfig->defaults->actions !== null) ? 'actions: '.\var_export($input->spec->fieldConfig->defaults->actions, true).', ' : '').''.(($input->spec->fieldConfig->defaults->noValue !== null) ? 'noValue: '.\var_export($input->spec->fieldConfig->defaults->noValue, true).', ' : '').''.(($input->spec->fieldConfig->defaults->custom !== null) ? 'custom: '.\var_export($input->spec->fieldConfig->defaults->custom, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if (count($input->spec->fieldConfig->overrides) >= 1) {
    
        
    $buffer = 'overrides(';
        $tmparg0 = [];
        foreach ($input->spec->fieldConfig->overrides as $arg1) {
        $tmpoverridesarg1 = \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1FieldConfigSourceOverridesConverter::convert($arg1);
        $tmparg0[] = $tmpoverridesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->spec->fieldConfig->overrides) >= 1) {
    foreach ($input->spec->fieldConfig->overrides as $item) {
        
    $buffer = 'overrideByName(';
        $arg0 ='(new \Grafana\Foundation\Dashboardv2beta1\MatcherConfig(id: '.\var_export($item->matcher->id, true).','.(($item->matcher->options !== null) ? 'options: '.\var_export($item->matcher->options, true).', ' : '').'))';
        $buffer .= $arg0;
        $buffer .= ', ';
        $tmparg1 = [];
        foreach ($item->properties as $arg1) {
        $tmppropertiesarg1 ='(new \Grafana\Foundation\Dashboardv2beta1\DynamicConfigValue(id: '.\var_export($arg1->id, true).','.(($arg1->value !== null) ? 'value: '.\var_export($arg1->value, true).', ' : '').'))';
        $tmparg1[] = $tmppropertiesarg1;
        }
        $arg1 = "[" . implode(", \n", $tmparg1) . "]";
        $buffer .= $arg1;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    }
    }

        return \implode("\n\t->", $calls);
    }
}

