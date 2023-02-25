import {MockoonServerless} from '@mockoon/serverless'
import mockEnv from "./Spotify.json" assert { type: "json" }

const mockoonServerless = new MockoonServerless(mockEnv)

export const handler = mockoonServerless.awsHandler()
