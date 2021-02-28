import type { NextApiRequest, NextApiResponse } from 'next'
import type { ServerEnv } from 'utils/env'
import { getEnvToUseServer } from 'utils/env'


export default async (_: NextApiRequest, res: NextApiResponse<ServerEnv>) => {
  const serverEnv = await getEnvToUseServer()
  res.status(200).json(serverEnv)
}