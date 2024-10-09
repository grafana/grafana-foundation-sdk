<?php

namespace Grafana\Foundation\Testdata;

final class CSVWaveConverter
{
    public static function convert(\Grafana\Foundation\Testdata\CSVWave $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Testdata\CSVWaveBuilder())',
        ];
            if ($input->labels !== null && $input->labels !== "") {
    
        
    $buffer = 'labels(';
        $arg0 =\var_export($input->labels, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->name !== null && $input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->timeStep !== null) {
    
        
    $buffer = 'timeStep(';
        $arg0 =\var_export($input->timeStep, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->valuesCSV !== null && $input->valuesCSV !== "") {
    
        
    $buffer = 'valuesCSV(';
        $arg0 =\var_export($input->valuesCSV, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

