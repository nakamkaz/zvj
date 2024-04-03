string filter command like as grep and awk

Usage: 
````
$ zvj 'REGEXP' < inputfile
````

````
$ command | zvj 'REGEXP'
````

For Windows :

````
zvj.exe "REGEXP" < inputfile
````
````
 command | zvj.exe 'REGEXP'
````

Option: 
-v  : Invert Match option
      Default value: false 

