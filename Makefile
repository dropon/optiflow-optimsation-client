generate:
		openapi-generator-cli generate -g go -i https://api.myptv.com/meta/services/routeoptimization/optiflow/v1/openapi.json
		# then set the go.mod and repally these commits:
		# c9d0eafdb4bdeaf81886de7b0492e86472186271
		# f6e633dadca0d15bb402e384ed6c6feb10ebe6c6	
		# e60f99ed41bff620b9babfaef3a93c6362cd0ec4
