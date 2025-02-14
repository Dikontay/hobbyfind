DIRS='./cmd/app,'
ENDPOINTS_DIR='./internal/endpoints/http'
for dir in "$ENDPOINTS_DIR"/*; do
    if [ -d "$dir" ]; then
        DIRS+="$dir,"
    fi
done
DIRS=${DIRS%,}
echo $DIRS
swag init --parseDependency  -d $DIRS -o ./cmd/docs
