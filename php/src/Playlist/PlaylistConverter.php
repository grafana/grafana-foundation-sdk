<?php

namespace Grafana\Foundation\Playlist;

final class PlaylistConverter
{
    public static function convert(\Grafana\Foundation\Playlist\Playlist $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Playlist\PlaylistBuilder())',
        ];
            if ($input->uid !== "") {
    
        
    $buffer = 'uid(';
        $arg0 =\var_export($input->uid, true);
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
            if ($input->interval !== "" && $input->interval !== "5m") {
    
        
    $buffer = 'interval(';
        $arg0 =\var_export($input->interval, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->items !== null && count($input->items) >= 1) {
    
        
    $buffer = 'items(';
        $tmparg0 = [];
        foreach ($input->items as $arg1) {
        $tmpitemsarg1 = \Grafana\Foundation\Playlist\PlaylistItemConverter::convert($arg1);
        $tmparg0[] = $tmpitemsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

