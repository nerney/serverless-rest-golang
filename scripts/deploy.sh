set -e 
echo "Deploying..."
serverless deploy
rm -rf ./bin