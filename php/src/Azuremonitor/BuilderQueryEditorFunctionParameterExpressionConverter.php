<?php

namespace Grafana\Foundation\Azuremonitor;

final class BuilderQueryEditorFunctionParameterExpressionConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\BuilderQueryEditorFunctionParameterExpression $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\BuilderQueryEditorFunctionParameterExpressionBuilder())',
        ];
            if ($input->value !== "") {
    
        
    $buffer = 'value(';
        $arg0 =\var_export($input->value, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'fieldType(';
        $arg0 ='\Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyType::fromValue("'.$input->fieldType.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType::fromValue("'.$input->type.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

