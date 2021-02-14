import { NextApiRequest, NextApiResponse } from 'next'

export const putHandler = (req: NextApiRequest, res: NextApiResponse) => {
  try {
    console.log(req)

    res.status(200).json({})
  } catch (err) {
    res.status(500).json({ statusCode: 500, message: err.message })
  }
}
