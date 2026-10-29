<?php

namespace Grafana\Foundation\Alertingv0alpha1;

final class FacetValueConverter
{
    public static function convert(\Grafana\Foundation\Alertingv0alpha1\FacetValue $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv0alpha1\FacetValueBuilder())',
        ];
            if ($input->value !== "") {
    
        
    $buffer = 'value(';
        $arg0 =\var_export($input->value, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'count(';
        $arg0 =\var_export($input->count, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

