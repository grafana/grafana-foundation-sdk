<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class ConditionalRenderingGroupConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupBuilder())',
        ];
            
    
        {
    $buffer = 'spec(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'visibility(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecVisibility::fromValue("'.$input->spec->visibility.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'condition(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecCondition::fromValue("'.$input->spec->condition.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if (count($input->spec->items) >= 1) {
    
        
    $buffer = 'items(';
        $tmparg0 = [];
        foreach ($input->spec->items as $arg1) {
        switch (true) {
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableConverter::convert($arg1);
                $tmpitemsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataConverter::convert($arg1);
                $tmpitemsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingTimeRangeSizeKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingTimeRangeSizeConverter::convert($arg1);
                $tmpitemsarg1 = $disjunctionarg1;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $tmparg0[] = $tmpitemsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

