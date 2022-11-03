#! /bin/bash

echo "1. Unzipping the file"
unzip -p one.zip > first.csv
unzip -p two.zip > second.csv

echo "2. Removing two redundant lines"
sed -e '2,/ImportId/d' first.csv > ${3}
sed -e '2,/ImportId/d' second.csv > ${4}

./invitation_list.r $1 $2 ${3} ${4}

rm first.csv
rm second.csv
rm ${3}
rm ${4}