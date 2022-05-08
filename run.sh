file="data/pantry.db"

if [ ! -e "$file" ] ; then
    touch "$file"
fi

if [ ! -w "$file" ] ; then
    echo cannot write to $file
    exit 1
fi

go run main.go