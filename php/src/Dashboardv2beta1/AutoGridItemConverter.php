<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class AutoGridItemConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutItemKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\AutoGridItemBuilder())',
        ];
            if ($input->spec->repeat !== null) {
    
        
    $buffer = 'repeat(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\AutoGridRepeatOptionsConverter::convert($input->spec->repeat);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->conditionalRendering !== null) {
    
        
    $buffer = 'conditionalRendering(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupConverter::convert($input->spec->conditionalRendering);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->element->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->spec->element->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

