<?php

namespace Grafana\Foundation\Dashboard;

final class AnnotationTargetConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\AnnotationTarget $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\AnnotationTargetBuilder())',
        ];
            
    
        {
    $buffer = 'limit(';
        $arg0 =\var_export($input->limit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'matchAny(';
        $arg0 =\var_export($input->matchAny, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if (count($input->tags) >= 1) {
    
        
    $buffer = 'tags(';
        $tmparg0 = [];
        foreach ($input->tags as $arg1) {
        $tmptagsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmptagsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->type !== "") {
    
        
    $buffer = 'type(';
        $arg0 =\var_export($input->type, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

