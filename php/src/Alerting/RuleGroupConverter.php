<?php

namespace Grafana\Foundation\Alerting;

final class RuleGroupConverter
{
    public static function convert(\Grafana\Foundation\Alerting\RuleGroup $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alerting\RuleGroupBuilder('.\var_export($input->title, true).'))',
        ];
            if ($input->folderUid !== null && $input->folderUid !== "") {
    
        
    $buffer = 'folderUid(';
        $arg0 =\var_export($input->folderUid, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'interval(';
        $arg0 =\var_export($input->interval, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->rules !== null && count($input->rules) >= 1) {
    
        
    $buffer = 'rules(';
        $tmparg0 = [];
        foreach ($input->rules as $arg1) {
        $tmprulesarg1 = \Grafana\Foundation\Alerting\RuleConverter::convert($arg1);
        $tmparg0[] = $tmprulesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->title !== null && $input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

