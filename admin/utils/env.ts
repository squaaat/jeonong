import AWS from 'aws-sdk'
import yaml from 'js-yaml'

type Env = {
  Env: string;
  Debug?: boolean;
  Project?: string;
  AppName?: string;

  RestHTTPServer?: RestHTTPServer;
}
type RestHTTPServer = {
  Url: string;
}

const env: Env = {
  Env: '',
}

// get call env using simple single tone variables
export default (): Promise<Env> => new Promise(async (resolve, reject) => {
  if (process.env.J_ENV === '' || process.env.J_ENV === undefined) {
    reject(new Error(`process.env.J_ENV is not defined actual value: [${process.env.J_ENV}]`))
  }
  if (process.env.J_CICD === '' || process.env.J_CICD === undefined) {
    reject(new Error(`process.env.J_CICD is not defined actual value: [${process.env.J_CICD}]`))
  }
  const ssmPath = `/nearsfeed/nearsfeed-admin/${process.env.J_ENV}/application.yml`

  if (env.Env === "") {
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

      initEnv(yml)
      resolve(env);
    });
  } else {
    resolve(env)
  }
})

const initEnv = (data: any) => {
  env.Env = data.env?.app?.env || ''
  env.Debug = data.env?.app?.debug || true
  env.Project = data.env?.app?.project || ''
  env.AppName = data.env?.app.app_name || ''

  env.RestHTTPServer = {
    Url: data.env?.rest_http_server?.url || ''
  }
}
