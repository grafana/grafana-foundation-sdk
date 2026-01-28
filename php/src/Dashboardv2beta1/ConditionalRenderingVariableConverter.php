<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class ConditionalRenderingVariableConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableBuilder())',
        ];
            
    
        {
    $buffer = 'spec(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpecConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec->variable !== "") {
    
        
    $buffer = 'variable(';
        $arg0 =\var_export($input->spec->variable, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'operator(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpecOperator::fromValue("'.$input->spec->operator.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec->value !== "") {
    
        
    $buffer = 'value(';
        $arg0 =\var_export($input->spec->value, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

