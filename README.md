### Parse vcf contacts file to csv

Go to icloud, export your contacts as vcf file.
Then run this script to convert it to csv file.

```
go run main.go
```

Or execute the binary `vcf_to_csv` (compiled on ARM mac) and it will prompt you for a path to the
vcf file, and a path to the csv file to export. Must use fully qualified file paths, 
Ie, `/Users/x/` instead of `~/`
