set -e 
echo "Running Tests..."
export GO111MODULE=on 
export TESTING=* 
echo "mode: set" > coverage.txt
# grepping out models and main since they don't have/need unit tests
for pkg in $( go list ./... | grep -v models | grep -v lambdas ); do
    go test -v -coverprofile=tmp.txt $pkg
    cat tmp.txt | tail -n +2 >> coverage.txt && rm tmp.txt
done
go tool cover -html=coverage.txt -o "coverage.html"

