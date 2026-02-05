<?php

namespace Grafana\Foundation\Logs;

final class VisualizationConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\VizConfigKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Logs\VisualizationBuilder())',
        ];
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
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options) {
    
        
    $buffer = 'showLabels(';
        $arg0 =\var_export($input->spec->options->showLabels, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options) {
    
        
    $buffer = 'showCommonLabels(';
        $arg0 =\var_export($input->spec->options->showCommonLabels, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options) {
    
        
    $buffer = 'showTime(';
        $arg0 =\var_export($input->spec->options->showTime, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options) {
    
        
    $buffer = 'showLogContextToggle(';
        $arg0 =\var_export($input->spec->options->showLogContextToggle, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options) {
    
        
    $buffer = 'wrapLogMessage(';
        $arg0 =\var_export($input->spec->options->wrapLogMessage, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options) {
    
        
    $buffer = 'prettifyLogMessage(';
        $arg0 =\var_export($input->spec->options->prettifyLogMessage, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options) {
    
        
    $buffer = 'enableLogDetails(';
        $arg0 =\var_export($input->spec->options->enableLogDetails, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options) {
    
        
    $buffer = 'sortOrder(';
        $arg0 ='\Grafana\Foundation\Common\LogsSortOrder::fromValue("'.$input->spec->options->sortOrder.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options) {
    
        
    $buffer = 'dedupStrategy(';
        $arg0 ='\Grafana\Foundation\Common\LogsDedupStrategy::fromValue("'.$input->spec->options->dedupStrategy.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options && $input->spec->options->enableInfiniteScrolling !== null) {
    
        
    $buffer = 'enableInfiniteScrolling(';
        $arg0 =\var_export($input->spec->options->enableInfiniteScrolling, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options && $input->spec->options->onClickFilterLabel !== null) {
    
        
    $buffer = 'onClickFilterLabel(';
        $arg0 =\var_export($input->spec->options->onClickFilterLabel, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options && $input->spec->options->onClickFilterOutLabel !== null) {
    
        
    $buffer = 'onClickFilterOutLabel(';
        $arg0 =\var_export($input->spec->options->onClickFilterOutLabel, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options && $input->spec->options->isFilterLabelActive !== null) {
    
        
    $buffer = 'isFilterLabelActive(';
        $arg0 =\var_export($input->spec->options->isFilterLabelActive, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options && $input->spec->options->onClickFilterString !== null) {
    
        
    $buffer = 'onClickFilterString(';
        $arg0 =\var_export($input->spec->options->onClickFilterString, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options && $input->spec->options->onClickFilterOutString !== null) {
    
        
    $buffer = 'onClickFilterOutString(';
        $arg0 =\var_export($input->spec->options->onClickFilterOutString, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options && $input->spec->options->onClickShowField !== null) {
    
        
    $buffer = 'onClickShowField(';
        $arg0 =\var_export($input->spec->options->onClickShowField, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options && $input->spec->options->onClickHideField !== null) {
    
        
    $buffer = 'onClickHideField(';
        $arg0 =\var_export($input->spec->options->onClickHideField, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options && $input->spec->options->logRowMenuIconsBefore !== null) {
    
        
    $buffer = 'logRowMenuIconsBefore(';
        $arg0 =\var_export($input->spec->options->logRowMenuIconsBefore, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options && $input->spec->options->logRowMenuIconsAfter !== null) {
    
        
    $buffer = 'logRowMenuIconsAfter(';
        $arg0 =\var_export($input->spec->options->logRowMenuIconsAfter, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options && $input->spec->options->onNewLogsReceived !== null) {
    
        
    $buffer = 'onNewLogsReceived(';
        $arg0 =\var_export($input->spec->options->onNewLogsReceived, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Logs\Options && $input->spec->options->displayedFields !== null && count($input->spec->options->displayedFields) >= 1) {
    
        
    $buffer = 'displayedFields(';
        $tmparg0 = [];
        foreach ($input->spec->options->displayedFields as $arg1) {
        $tmpdisplayedFieldsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpdisplayedFieldsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

