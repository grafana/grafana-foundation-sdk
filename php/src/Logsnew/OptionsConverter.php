<?php

namespace Grafana\Foundation\Logsnew;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Logsnew\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Logsnew\OptionsBuilder())',
        ];
            
    
        {
    $buffer = 'showTime(';
        $arg0 =\var_export($input->showTime, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'wrapLogMessage(';
        $arg0 =\var_export($input->wrapLogMessage, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'enableLogDetails(';
        $arg0 =\var_export($input->enableLogDetails, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'sortOrder(';
        $arg0 ='\Grafana\Foundation\Common\LogsSortOrder::fromValue("'.$input->sortOrder.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'dedupStrategy(';
        $arg0 ='\Grafana\Foundation\Common\LogsDedupStrategy::fromValue("'.$input->dedupStrategy.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->enableInfiniteScrolling !== null) {
    
        
    $buffer = 'enableInfiniteScrolling(';
        $arg0 =\var_export($input->enableInfiniteScrolling, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->onNewLogsReceived !== null) {
    
        
    $buffer = 'onNewLogsReceived(';
        $arg0 =\var_export($input->onNewLogsReceived, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

