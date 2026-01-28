<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class AutoGridLayoutItemSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutItemSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutItemSpecBuilder())',
        ];
            
    
        {
    $buffer = 'element(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\ElementReferenceConverter::convert($input->element);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->repeat !== null) {
    
        
    $buffer = 'repeat(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\AutoGridRepeatOptionsConverter::convert($input->repeat);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->conditionalRendering !== null) {
    
        
    $buffer = 'conditionalRendering(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupConverter::convert($input->conditionalRendering);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

