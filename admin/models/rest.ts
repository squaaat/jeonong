import axios, { AxiosInstance } from 'axios'
import { getEnvToUseClient, getEnvToUseServer } from 'utils/env'
import isBrowser from 'utils/isBrowser'

let restAxios: AxiosInstance

export const RestAxios = async ():Promise<AxiosInstance> => {
  if (isBrowser()) {
    return RestHTTPClientAxios()
  } else {
    return RestHTTPServerAxios()
  }
}

export const RestHTTPClientAxios = async ():Promise<AxiosInstance> => {
  if(!restAxios) {
    const e = await getEnvToUseClient()
    restAxios = axios.create({
      baseURL: e.RestHTTPServer?.Url,
    });
  }
  return restAxios
}

export const RestHTTPServerAxios = async ():Promise<AxiosInstance> => {
  if(!restAxios) {
    const e = await getEnvToUseServer()
    restAxios = axios.create({
      baseURL: e.RestHTTPServer?.Url,
    });
  }
  return restAxios
}
