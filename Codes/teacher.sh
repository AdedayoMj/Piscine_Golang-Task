#!/bin/bash
mystery=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d "#" -f2)
echo $mystery
cat interviews/interview-$mystery
grep -A 5  "L337" vehicles  | grep -A 4 -B 1 "Honda" | grep -A 3 -B 2 "Blue" | grep -B 4 "6'" 
cat memberships/AAA memberships/Delta_SkyMiles memberships/Museum_of_Bash_History memberships/Terminal_City_Library | grep "$MAIN_SUSPECT" | wc -l
# cat memberships/AAA memberships/Delta_SkyMiles memberships/Museum_of_Bash_History memberships/Terminal_City_Library |grep -c "$MAIN_SUSPECT"