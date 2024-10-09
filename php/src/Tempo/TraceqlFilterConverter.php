<?php

namespace Grafana\Foundation\Tempo;

final class TraceqlFilterConverter
{
    public static function convert(\Grafana\Foundation\Tempo\TraceqlFilter $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Tempo\TraceqlFilterBuilder())',
        ];
            if ($input->id !== "") {
    
        
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->tag !== null && $input->tag !== "") {
    
        
    $buffer = 'tag(';
        $arg0 =\var_export($input->tag, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->operator !== null && $input->operator !== "") {
    
        
    $buffer = 'operator(';
        $arg0 =\var_export($input->operator, true);
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
            case is_array($input->value):
                $tmpdisjunctionvalue = [];
        foreach ($input->value as $arg1) {
        $tmpvaluearg1 =\var_export($arg1, true);
        $tmpdisjunctionvalue[] = $tmpvaluearg1;
        }
        $disjunctionvalue = "[" . implode(", \n", $tmpdisjunctionvalue) . "]";
                $arg0 = $disjunctionvalue;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->valueType !== null && $input->valueType !== "") {
    
        
    $buffer = 'valueType(';
        $arg0 =\var_export($input->valueType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->scope !== null) {
    
        
    $buffer = 'scope(';
        $arg0 ='\Grafana\Foundation\Tempo\TraceqlSearchScope::fromValue("'.$input->scope.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

