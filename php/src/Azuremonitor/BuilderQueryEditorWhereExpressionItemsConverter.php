<?php

namespace Grafana\Foundation\Azuremonitor;

final class BuilderQueryEditorWhereExpressionItemsConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionItems $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionItemsBuilder())',
        ];
            
    
        {
    $buffer = 'property(';
        $arg0 = \Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyConverter::convert($input->property);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'operator(';
        $arg0 = \Grafana\Foundation\Azuremonitor\BuilderQueryEditorOperatorConverter::convert($input->operator);
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

