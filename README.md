# ratio

Generate differents ratios of an image.

## Installation

```
$ git clone https://github.com/eze-kiel/ratio.git
$ cd ratio
$ go build
```

## Usage

```
Usage: ratio [-i] [width] [height]
```

Where `-i` is the increase flag.

Examples:

```
$ ratio 1920 1080
%	width	height
10	192	108
20	384	216
30	576	324
40	768	432
50	960	540
60	1152	648
70	1344	756
80	1536	864
90	1728	972
100	1920	1080
```

With the increase flag (`-i`):

```
%       width   height
110     2112    1188
120     2304    1296
130     2496    1404
140     2688    1512
150     2880    1620
160     3072    1728
170     3264    1836
180     3456    1944
190     3648    2052
200     3840    2160
```