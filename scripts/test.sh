set -e 
echo "Running Tests..."
export GO111MODULE=on 
export TESTING=* 
echo "mode: set" > coverage.txt
for pkg in $( go list ./... | tail -n +2 | grep -v models | grep -v lambda | grep -v storage ); do
    go test -coverprofile=tmp.txt $pkg
    cat tmp.txt | tail -n +2 >> coverage.txt
done
go tool cover -html=coverage.txt -o "coverage.html"
rm tmp.txt

