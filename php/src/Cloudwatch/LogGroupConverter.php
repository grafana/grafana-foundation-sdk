<?php

namespace Grafana\Foundation\Cloudwatch;

final class LogGroupConverter
{
    public static function convert(\Grafana\Foundation\Cloudwatch\LogGroup $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Cloudwatch\LogGroupBuilder())',
        ];
            if ($input->arn !== "") {
    
        
    $buffer = 'arn(';
        $arg0 =\var_export($input->arn, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->accountId !== null && $input->accountId !== "") {
    
        
    $buffer = 'accountId(';
        $arg0 =\var_export($input->accountId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->accountLabel !== null && $input->accountLabel !== "") {
    
        
    $buffer = 'accountLabel(';
        $arg0 =\var_export($input->accountLabel, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

