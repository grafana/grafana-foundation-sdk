<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class MetricFindValueConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\MetricFindValue $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\MetricFindValueBuilder())',
        ];
            if ($input->text !== "") {
    
        
    $buffer = 'text(';
        $arg0 =\var_export($input->text, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->value !== null) {
    
        
    $buffer = 'value(';
        switch (true) {
            case is_string($input->value):
                $disjunctionvalue =\var_export($input->value, true);
                $arg0 = $disjunctionvalue;
                break;
            case is_float($input->value):
                $disjunctionvalue =\var_export($input->value, true);
                $arg0 = $disjunctionvalue;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->group !== null && $input->group !== "") {
    
        
    $buffer = 'group(';
        $arg0 =\var_export($input->group, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->expandable !== null) {
    
        
    $buffer = 'expandable(';
        $arg0 =\var_export($input->expandable, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

