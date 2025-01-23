<?php

namespace Grafana\Foundation\Stat;

final class PanelConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\Panel $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Stat\PanelBuilder())',
        ];
            if ($input->id !== null) {
    
        
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->targets !== null && count($input->targets) >= 1) {
    
        
    $buffer = 'targets(';
        $tmparg0 = [];
        foreach ($input->targets as $arg1) {
        $tmptargetsarg1 = \Grafana\Foundation\Cog\Runtime::get()->convertDataqueryToCode($arg1, );
        $tmparg0[] = $tmptargetsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->title !== null && $input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->description !== null && $input->description !== "") {
    
        
    $buffer = 'description(';
        $arg0 =\var_export($input->description, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->transparent !== null && $input->transparent !== false) {
    
        
    $buffer = 'transparent(';
        $arg0 =\var_export($input->transparent, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\DataSourceRef('.(($input->datasource->type !== null) ? 'type: '.\var_export($input->datasource->type, true).', ' : '').''.(($input->datasource->uid !== null) ? 'uid: '.\var_export($input->datasource->uid, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->gridPos !== null) {
    
        
    $buffer = 'gridPos(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\GridPos(h: '.\var_export($input->gridPos->h, true).',w: '.\var_export($input->gridPos->w, true).',x: '.\var_export($input->gridPos->x, true).',y: '.\var_export($input->gridPos->y, true).','.(($input->gridPos->static !== null) ? 'static: '.\var_export($input->gridPos->static, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->gridPos !== null && $input->gridPos->h !== 9) {
    
        
    $buffer = 'height(';
        $arg0 =\var_export($input->gridPos->h, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->gridPos !== null && $input->gridPos->w !== 12) {
    
        
    $buffer = 'span(';
        $arg0 =\var_export($input->gridPos->w, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->links !== null && count($input->links) >= 1) {
    
        
    $buffer = 'links(';
        $tmparg0 = [];
        foreach ($input->links as $arg1) {
        $tmplinksarg1 = \Grafana\Foundation\Dashboard\DashboardLinkConverter::convert($arg1);
        $tmparg0[] = $tmplinksarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->repeat !== null && $input->repeat !== "") {
    
        
    $buffer = 'repeat(';
        $arg0 =\var_export($input->repeat, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->repeatDirection !== null) {
    
        
    $buffer = 'repeatDirection(';
        $arg0 ='\Grafana\Foundation\Dashboard\PanelRepeatDirection::fromValue("'.$input->repeatDirection.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->maxPerRow !== null) {
    
        
    $buffer = 'maxPerRow(';
        $arg0 =\var_export($input->maxPerRow, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->maxDataPoints !== null) {
    
        
    $buffer = 'maxDataPoints(';
        $arg0 =\var_export($input->maxDataPoints, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->transformations !== null && count($input->transformations) >= 1) {
    
        
    $buffer = 'transformations(';
        $tmparg0 = [];
        foreach ($input->transformations as $arg1) {
        $tmptransformationsarg1 ='(new \Grafana\Foundation\Dashboard\DataTransformerConfig(id: '.\var_export($arg1->id, true).','.(($arg1->disabled !== null) ? 'disabled: '.\var_export($arg1->disabled, true).', ' : '').''.(($arg1->filter !== null) ? 'filter: '.'(new \Grafana\Foundation\Dashboard\MatcherConfig(id: '.\var_export($arg1->filter->id, true).','.(($arg1->filter->options !== null) ? 'options: '.\var_export($arg1->filter->options, true).', ' : '').'))'.', ' : '').''.(($arg1->topic !== null) ? 'topic: '.'\Grafana\Foundation\Dashboard\DataTransformerConfigTopic::fromValue("'.$arg1->topic.'")'.', ' : '').'options: '.\var_export($arg1->options, true).',))';
        $tmparg0[] = $tmptransformationsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->interval !== null && $input->interval !== "") {
    
        
    $buffer = 'interval(';
        $arg0 =\var_export($input->interval, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->timeFrom !== null && $input->timeFrom !== "") {
    
        
    $buffer = 'timeFrom(';
        $arg0 =\var_export($input->timeFrom, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->timeShift !== null && $input->timeShift !== "") {
    
        
    $buffer = 'timeShift(';
        $arg0 =\var_export($input->timeShift, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->hideTimeOverride !== null) {
    
        
    $buffer = 'hideTimeOverride(';
        $arg0 =\var_export($input->hideTimeOverride, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->libraryPanel !== null) {
    
        
    $buffer = 'libraryPanel(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\LibraryPanelRef(name: '.\var_export($input->libraryPanel->name, true).',uid: '.\var_export($input->libraryPanel->uid, true).',))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->displayName !== null && $input->fieldConfig->defaults->displayName !== "") {
    
        
    $buffer = 'displayName(';
        $arg0 =\var_export($input->fieldConfig->defaults->displayName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->unit !== null && $input->fieldConfig->defaults->unit !== "") {
    
        
    $buffer = 'unit(';
        $arg0 =\var_export($input->fieldConfig->defaults->unit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->decimals !== null) {
    
        
    $buffer = 'decimals(';
        $arg0 =\var_export($input->fieldConfig->defaults->decimals, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->min !== null) {
    
        
    $buffer = 'min(';
        $arg0 =\var_export($input->fieldConfig->defaults->min, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->max !== null) {
    
        
    $buffer = 'max(';
        $arg0 =\var_export($input->fieldConfig->defaults->max, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->mappings !== null && count($input->fieldConfig->defaults->mappings) >= 1) {
    
        
    $buffer = 'mappings(';
        $tmparg0 = [];
        foreach ($input->fieldConfig->defaults->mappings as $arg1) {
        switch (true) {
            case $arg1 instanceof \Grafana\Foundation\Dashboard\ValueMap:
                $disjunctionarg1 ='(new \Grafana\Foundation\Dashboard\ValueMap(type: '.\var_export($arg1->type, true).',options: '.\var_export($arg1->options, true).',))';
                $tmpmappingsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboard\RangeMap:
                $disjunctionarg1 ='(new \Grafana\Foundation\Dashboard\RangeMap(type: '.\var_export($arg1->type, true).',options: '.'(new \Grafana\Foundation\Dashboard\DashboardRangeMapOptions('.(($arg1->options->from !== null) ? 'from: '.\var_export($arg1->options->from, true).', ' : '').''.(($arg1->options->to !== null) ? 'to: '.\var_export($arg1->options->to, true).', ' : '').'result: '.'(new \Grafana\Foundation\Dashboard\ValueMappingResult('.(($arg1->options->result->text !== null) ? 'text: '.\var_export($arg1->options->result->text, true).', ' : '').''.(($arg1->options->result->color !== null) ? 'color: '.\var_export($arg1->options->result->color, true).', ' : '').''.(($arg1->options->result->icon !== null) ? 'icon: '.\var_export($arg1->options->result->icon, true).', ' : '').''.(($arg1->options->result->index !== null) ? 'index: '.\var_export($arg1->options->result->index, true).', ' : '').'))'.',))'.',))';
                $tmpmappingsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboard\RegexMap:
                $disjunctionarg1 ='(new \Grafana\Foundation\Dashboard\RegexMap(type: '.\var_export($arg1->type, true).',options: '.'(new \Grafana\Foundation\Dashboard\DashboardRegexMapOptions(pattern: '.\var_export($arg1->options->pattern, true).',result: '.'(new \Grafana\Foundation\Dashboard\ValueMappingResult('.(($arg1->options->result->text !== null) ? 'text: '.\var_export($arg1->options->result->text, true).', ' : '').''.(($arg1->options->result->color !== null) ? 'color: '.\var_export($arg1->options->result->color, true).', ' : '').''.(($arg1->options->result->icon !== null) ? 'icon: '.\var_export($arg1->options->result->icon, true).', ' : '').''.(($arg1->options->result->index !== null) ? 'index: '.\var_export($arg1->options->result->index, true).', ' : '').'))'.',))'.',))';
                $tmpmappingsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboard\SpecialValueMap:
                $disjunctionarg1 ='(new \Grafana\Foundation\Dashboard\SpecialValueMap(type: '.\var_export($arg1->type, true).',options: '.'(new \Grafana\Foundation\Dashboard\DashboardSpecialValueMapOptions(match: '.'\Grafana\Foundation\Dashboard\SpecialValueMatch::fromValue("'.$arg1->options->match.'")'.',result: '.'(new \Grafana\Foundation\Dashboard\ValueMappingResult('.(($arg1->options->result->text !== null) ? 'text: '.\var_export($arg1->options->result->text, true).', ' : '').''.(($arg1->options->result->color !== null) ? 'color: '.\var_export($arg1->options->result->color, true).', ' : '').''.(($arg1->options->result->icon !== null) ? 'icon: '.\var_export($arg1->options->result->icon, true).', ' : '').''.(($arg1->options->result->index !== null) ? 'index: '.\var_export($arg1->options->result->index, true).', ' : '').'))'.',))'.',))';
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
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->thresholds !== null) {
    
        
    $buffer = 'thresholds(';
        $arg0 = \Grafana\Foundation\Dashboard\ThresholdsConfigConverter::convert($input->fieldConfig->defaults->thresholds);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->color !== null) {
    
        
    $buffer = 'colorScheme(';
        $arg0 = \Grafana\Foundation\Dashboard\FieldColorConverter::convert($input->fieldConfig->defaults->color);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->links !== null && count($input->fieldConfig->defaults->links) >= 1) {
    
        
    $buffer = 'dataLinks(';
        $tmparg0 = [];
        foreach ($input->fieldConfig->defaults->links as $arg1) {
        $tmplinksarg1 = \Grafana\Foundation\Dashboard\DashboardLinkConverter::convert($arg1);
        $tmparg0[] = $tmplinksarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->noValue !== null && $input->fieldConfig->defaults->noValue !== "") {
    
        
    $buffer = 'noValue(';
        $arg0 =\var_export($input->fieldConfig->defaults->noValue, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && count($input->fieldConfig->overrides) >= 1) {
    
        
    $buffer = 'overrides(';
        $tmparg0 = [];
        foreach ($input->fieldConfig->overrides as $arg1) {
        $tmpoverridesarg1 = \Grafana\Foundation\Dashboard\DashboardFieldConfigSourceOverridesConverter::convert($arg1);
        $tmparg0[] = $tmpoverridesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && count($input->fieldConfig->overrides) >= 1) {
    foreach ($input->fieldConfig->overrides as $item) {
        
    $buffer = 'withOverride(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\MatcherConfig(id: '.\var_export($item->matcher->id, true).','.(($item->matcher->options !== null) ? 'options: '.\var_export($item->matcher->options, true).', ' : '').'))';
        $buffer .= $arg0;
        $buffer .= ', ';
        $tmparg1 = [];
        foreach ($item->properties as $arg1) {
        $tmppropertiesarg1 ='(new \Grafana\Foundation\Dashboard\DynamicConfigValue(id: '.\var_export($arg1->id, true).','.(($arg1->value !== null) ? 'value: '.\var_export($arg1->value, true).', ' : '').'))';
        $tmparg1[] = $tmppropertiesarg1;
        }
        $arg1 = "[" . implode(", \n", $tmparg1) . "]";
        $buffer .= $arg1;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    }
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Stat\Options) {
    
        
    $buffer = 'graphMode(';
        $arg0 ='\Grafana\Foundation\Common\BigValueGraphMode::fromValue("'.$input->options->graphMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Stat\Options) {
    
        
    $buffer = 'colorMode(';
        $arg0 ='\Grafana\Foundation\Common\BigValueColorMode::fromValue("'.$input->options->colorMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Stat\Options) {
    
        
    $buffer = 'justifyMode(';
        $arg0 ='\Grafana\Foundation\Common\BigValueJustifyMode::fromValue("'.$input->options->justifyMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Stat\Options) {
    
        
    $buffer = 'textMode(';
        $arg0 ='\Grafana\Foundation\Common\BigValueTextMode::fromValue("'.$input->options->textMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Stat\Options && $input->options->wideLayout !== true) {
    
        
    $buffer = 'wideLayout(';
        $arg0 =\var_export($input->options->wideLayout, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Stat\Options) {
    
        
    $buffer = 'reduceOptions(';
        $arg0 = \Grafana\Foundation\Common\ReduceDataOptionsConverter::convert($input->options->reduceOptions);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Stat\Options && $input->options->text !== null) {
    
        
    $buffer = 'text(';
        $arg0 = \Grafana\Foundation\Common\VizTextDisplayOptionsConverter::convert($input->options->text);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Stat\Options && $input->options->showPercentChange !== false) {
    
        
    $buffer = 'showPercentChange(';
        $arg0 =\var_export($input->options->showPercentChange, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Stat\Options) {
    
        
    $buffer = 'orientation(';
        $arg0 ='\Grafana\Foundation\Common\VizOrientation::fromValue("'.$input->options->orientation.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

