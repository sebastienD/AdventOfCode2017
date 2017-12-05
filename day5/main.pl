use strict;
use warnings;

my $file = 'input.txt';
open my $info, $file or die "Could not open $file: $!";

my @lines = ();

while( my $line = <$info>)  {   
    push(@lines, $line);
}

close $info;

#@lines = (0,3,0,1,-3);

my $i = 0;
my $steps = 0;

my $loop = 0;

my $size = @lines;

while (($i < @lines) && ($i >= 0)) {
    my $offset = $lines[$i];
    my $lasti = $i;
    $i += $offset;

    if ($offset >= 3) {
        $lines[$lasti]--;
    } else {
        $lines[$lasti]++;
    }

    #print "$lines[$lasti] index $i \n";
    $steps++;
    $loop++;
    last if ($loop > 30000000);
}

print "Steps $steps\n";

