<?php

namespace Grafana\Foundation\Azuremonitor;

final class BuilderQueryEditorOperatorConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\BuilderQueryEditorOperator $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\BuilderQueryEditorOperatorBuilder())',
        ];
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
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
            if ($input->labelValue !== null && $input->labelValue !== "") {
    
        
    $buffer = 'labelValue(';
        $arg0 =\var_export($input->labelValue, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

