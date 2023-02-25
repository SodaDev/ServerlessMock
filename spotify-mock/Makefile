build:
	sam build

deploy: build
	sam deploy \
		--stack-name=spotify-mock \
		--region=eu-west-1 \
		--resolve-s3 \
		--capabilities CAPABILITY_IAM

logs:
	sam logs -n SpotifyMockLambda --stack-name spotify-mock --tail

delete:
	aws cloudformation delete-stack --stack-name spotify-mock