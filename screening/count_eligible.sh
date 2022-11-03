#! /bin/bash

echo "1. Unzipping the file"
unzip -p one.zip > screen.csv

echo "2. Removing two redundant lines"
sed -e '2,/ImportId/d' screen.csv > screening.csv
rm screen.csv

total=$(awk -F "\"*,\"*" '{print $1}' screening.csv | awk -F " " '{print $1}' | awk -F "-" '{print $2,$3, $1}' | awk 'BEGIN { FIELDWIDTH = "4 2 2"} {printf "%s-%s-%s\n", $3, $1, $2}' | awk -v date=$1 '{if($1>=date ) print $1}' | wc -l)
echo "3. Counting the entries in the interval starting at $1: $total"

el=$(awk -vFPAT='([^,]*)|("[^"]+")' -vOFS=',' -v date=$1 '{if($1>=date && $5 >= 99 && $180 != "" && $184 >= 10 && $185 < 21) print $1, $5, $42, $43, $44, $45, $35, $180, $184, $185}' OFS=" " screening.csv | wc -l)
echo "4. Counting the number of eligible participants starting at $1: $el"

echo "5. Send to the manager: $el/$total"
if [[ $el -ge 120 ]]
then
	echo "6. You have enough eligible participants for the trial"
else
	echo "6. The number of eligible participants is still not enough"
fi

rm screening.csv
