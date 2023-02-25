import { handler } from '../../../src/handlers/spotify-mock-lambda.mjs';
import {jest} from '@jest/globals'
jest.setTimeout(10000)

describe('Test usage of serverless mock api', function () {
    it('Verifies successful response', async () => {
        const event = {
            "body": null,
            "resource": "/{proxy+}",
            "path": "/v1/albums",
            "httpMethod": "GET",
            "isBase64Encoded": false,
            "queryStringParameters": {},
            "multiValueQueryStringParameters": {},
            "pathParameters": {
                "proxy": "/v1/albums"
            },
            "stageVariables": {},
            "headers": {},
            "multiValueHeaders": {},
            "requestContext": {
                "accountId": "123456789012",
                "resourceId": "123456",
                "stage": "prod",
                "requestId": "c6af9ac6-7b61-11e6-9a41-93e8deadbeef",
                "requestTime": "09/Apr/2015:12:34:56 +0000",
                "requestTimeEpoch": 1428582896000,
                "identity": {
                    "sourceIp": "127.0.0.1",
                    "userAgent": "Custom User Agent String"
                },
                "path": "/v1/albums",
                "resourcePath": "/{proxy+}",
                "httpMethod": "POST",
                "apiId": "1234567890",
                "protocol": "HTTP/1.1"
            }
        }

        const result = await handler(event);

        expect(result.statusCode).toEqual(200)
    });

    it('Verifies throttled response', async () => {
        const event = {
            "body": null,
            "resource": "/{proxy+}",
            "path": "/v1/albums",
            "httpMethod": "GET",
            "isBase64Encoded": false,
            "queryStringParameters": {},
            "multiValueQueryStringParameters": {},
            "pathParameters": {
                "proxy": "/v1/albums"
            },
            "stageVariables": {},
            "headers": {},
            "multiValueHeaders": {
                "Scenario": ["Throttle"]
            },
            "requestContext": {
                "accountId": "123456789012",
                "resourceId": "123456",
                "stage": "prod",
                "requestId": "c6af9ac6-7b61-11e6-9a41-93e8deadbeef",
                "requestTime": "09/Apr/2015:12:34:56 +0000",
                "requestTimeEpoch": 1428582896000,
                "identity": {
                    "sourceIp": "127.0.0.1",
                    "userAgent": "Custom User Agent String"
                },
                "path": "/v1/albums",
                "resourcePath": "/{proxy+}",
                "httpMethod": "POST",
                "apiId": "1234567890",
                "protocol": "HTTP/1.1"
            }
        }

        const result = await handler(event);

        expect(result.statusCode).toEqual(429)
    });
});
