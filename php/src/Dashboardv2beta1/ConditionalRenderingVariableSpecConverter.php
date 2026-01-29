<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class ConditionalRenderingVariableSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpecBuilder())',
        ];
            if ($input->variable !== "") {
    
        
    $buffer = 'variable(';
        $arg0 =\var_export($input->variable, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'operator(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpecOperator::fromValue("'.$input->operator.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->value !== "") {
    
        
    $buffer = 'value(';
        $arg0 =\var_export($input->value, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

