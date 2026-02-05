<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class VariableValueOptionConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\VariableValueOption $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\VariableValueOptionBuilder())',
        ];
            if ($input->label !== "") {
    
        
    $buffer = 'label(';
        $arg0 =\var_export($input->label, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'value(';
        switch (true) {
            case is_string($input->value):
                $disjunctionvalue =\var_export($input->value, true);
                $arg0 = $disjunctionvalue;
                break;
            case is_bool($input->value):
                $disjunctionvalue =\var_export($input->value, true);
                $arg0 = $disjunctionvalue;
                break;
            case is_float($input->value):
                $disjunctionvalue =\var_export($input->value, true);
                $arg0 = $disjunctionvalue;
                break;
            case $input->value instanceof \Grafana\Foundation\Dashboardv2beta1\CustomVariableValue:
                $disjunctionvalue ='(new \Grafana\Foundation\Dashboardv2beta1\CustomVariableValue('.(($input->value->formatter !== null) ? 'formatter: '.\var_export($input->value->formatter, true).', ' : '').'))';
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

        return \implode("\n\t->", $calls);
    }
}

