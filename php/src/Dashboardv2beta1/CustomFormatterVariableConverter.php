<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class CustomFormatterVariableConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\CustomFormatterVariable $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\CustomFormatterVariableBuilder())',
        ];
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\VariableType::fromValue("'.$input->type.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'multi(';
        $arg0 =\var_export($input->multi, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'includeAll(';
        $arg0 =\var_export($input->includeAll, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

