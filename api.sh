function start_project(){
    go mod tidy
    echo "\
    +-----------------+
    | Starting api... |
    +-----------------+\
    "
    cd ./src/api
    go run ./main.go
    cd ../..
}

if [ -z "$1" ]; then
    echo "Usage: run.sh -run"
    return 1
fi

case "$1" in
    -run)
        start_project
        ;;
    *)
        echo "Invalid option. Usage: ./run.sh <-run>"
        return 1
        ;;
esac