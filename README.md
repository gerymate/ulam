![Ulam-spiral](https://github.com/user-attachments/assets/d9914458-0f3f-48bb-87e4-1f77ab0deae1)
# ulam
Tinkering with Ulam's arrangement of natural numbers. Highlighting primes.

## build and run
*$ go build*

or just
    
*$ go run ulam.go*

## command line options
*$ ulam -h*

## examples
*$ go run ulam.go -hlm mark -m1 \( -m2 \) -n 23*
```
(17)  16   15   14  (13)
 18  ( 5)   4  ( 3)  12 
(19)   6    1  ( 2) (11)
 20  ( 7)   8    9   10 
 21   22  (23)
```
*$ go run ulam.go -hlm eraseothers -n 40*
```
37                31
   17          13   
       5     3    29
   19        2 11   
       7            
         23         
```
*$ go run ulam.go -hlm erase*
```
145 144 143 142 141 140     138     136 135 134 133
146     100  99  98      96  95  94  93  92  91 132
147 102  65  64  63  62      60      58  57  90    
148      66      36  35  34  33  32      56     130
    104      38      16  15  14      30  55  88 129
150 105  68  39  18      4       12      54  87 128
    106  69  40      6   1           28      86    
         70      20      8   9   10  27  52  85 126
    108      42  21  22      24  25  26  51  84 125
         72      44  45  46      48  49  50     124
    110      74  75  76  77  78      80  81  82 123
    111 112     114 115 116 117 118 119 120 121 122
```