import AWS from 'aws-sdk'
import yaml from 'js-yaml'
import axios from 'axios'


export type ServerEnv = {
  Env: string;
  Debug?: boolean;
  Project?: string;
  AppName?: string;

  RestHTTPServer?: RestHTTPServer;
}
type RestHTTPServer = {
  Url: string;
}

const serverEnv: ServerEnv = {
  Env: '',
}

export type ClientEnv = {
  Env: string;
  Debug?: boolean;
  Project?: string;
  AppName?: string;

  RestHTTPServer?: RestHTTPServer;
}

const clientEnv: ClientEnv = {
  Env: '',
}


// get call env using simple single tone variables
export const getEnvToUseServer = (): Promise<ServerEnv> => new Promise(async (resolve, reject) => {
  if (process.env.J_ENV === '' || process.env.J_ENV === undefined) {
    reject(new Error(`process.env.J_ENV is not defined actual value: [${process.env.J_ENV}]`))
  }
  if (process.env.J_CICD === '' || process.env.J_CICD === undefined) {
    reject(new Error(`process.env.J_CICD is not defined actual value: [${process.env.J_CICD}]`))
  }
  const ssmPath = `/nearsfeed/nearsfeed-admin/${process.env.J_ENV}/application.yml`

  if (serverEnv.Env === "") {
    const ssm = new AWS.SSM()
    ssm.getParameter({
      Name: ssmPath,
      WithDecryption: true
    }, (err, data) => {
      if (err) reject(err);

      const subErr = new Error(`ssm path [${ssmPath}] has no value, ${data.Parameter?.Value}`)

      if (data.Parameter === undefined) reject(subErr)

      let yml: any = yaml.load(data.Parameter?.Value || '');
      if ((yml.env?.app?.env || '') === '') reject(subErr)

      initServerEnv(yml)
      resolve(serverEnv);
    });
  } else {
    resolve(serverEnv)
  }
})

// get call env using simple single tone variables
export const getEnvToUseClient = (): Promise<ClientEnv> => new Promise(async (resolve, reject) => {
  if (clientEnv.Env === "") {
    try {
      const res = await axios.get("/api/env")
      initClientEnv(res.data)
      resolve(clientEnv);
    } catch (e) {
      reject(e)
    }
  } else {
    resolve(clientEnv)
  }
})

const initServerEnv = (data: any) => {
  serverEnv.Env = data.env?.app?.env || ''
  serverEnv.Debug = data.env?.app?.debug || true
  serverEnv.Project = data.env?.app?.project || ''
  serverEnv.AppName = data.env?.app.app_name || ''

  serverEnv.RestHTTPServer = {
    Url: data.env?.rest_http_server?.url || ''
  }
}
const initClientEnv = (data: ClientEnv) => {
  if (!data) return
  if (!data.Env) return
  if (!data.Debug) return
  if (!data.Project) return
  if (!data.AppName) return
  if (!data.RestHTTPServer) return
  if (!data.RestHTTPServer.Url) return

  clientEnv.Env = data.Env
  clientEnv.Debug = data.Debug
  clientEnv.Project = data.Project
  clientEnv.AppName = data.AppName
  clientEnv.RestHTTPServer = data.RestHTTPServer
}
