<?php

namespace Grafana\Foundation\Azuremonitor;

final class BuilderQueryEditorOrderByExpressionConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\BuilderQueryEditorOrderByExpression $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\BuilderQueryEditorOrderByExpressionBuilder())',
        ];
            
    
        {
    $buffer = 'property(';
        $arg0 = \Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyConverter::convert($input->property);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'order(';
        $arg0 ='\Grafana\Foundation\Azuremonitor\BuilderQueryEditorOrderByOptions::fromValue("'.$input->order.'")';
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

