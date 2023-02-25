import {MockoonServerless} from '@mockoon/serverless'
import {GetObjectCommand, S3Client} from "@aws-sdk/client-s3";

const s3Client = new S3Client({region: 'eu-west-1'})
const s3Object = await s3Client.send(new GetObjectCommand({
    Key: "Spotify.json",
    Bucket: process.env.DEFINITIONS_BUCKET
}))
const mockEnv = JSON.parse(await s3Object.Body.transformToString())
const mockoonServerless = new MockoonServerless(mockEnv)

export const handler = mockoonServerless.awsHandler()
