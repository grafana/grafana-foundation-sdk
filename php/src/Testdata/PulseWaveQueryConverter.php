<?php

namespace Grafana\Foundation\Testdata;

final class PulseWaveQueryConverter
{
    public static function convert(\Grafana\Foundation\Testdata\PulseWaveQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Testdata\PulseWaveQueryBuilder())',
        ];
            if ($input->timeStep !== null) {
    
        
    $buffer = 'timeStep(';
        $arg0 =\var_export($input->timeStep, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->onCount !== null) {
    
        
    $buffer = 'onCount(';
        $arg0 =\var_export($input->onCount, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->offCount !== null) {
    
        
    $buffer = 'offCount(';
        $arg0 =\var_export($input->offCount, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->onValue !== null) {
    
        
    $buffer = 'onValue(';
        $arg0 =\var_export($input->onValue, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->offValue !== null) {
    
        
    $buffer = 'offValue(';
        $arg0 =\var_export($input->offValue, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

