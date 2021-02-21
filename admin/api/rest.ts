import axios, { AxiosInstance } from 'axios'
import env from 'utils/env'

type Result = {
  success: boolean;
  message?: string;
  error?: Error | object | any | undefined;
  GetCategoriesResult?: CategoryFromServer[];
}


export type CategoryFromServer = {
  ID: string;
  Status: string;
  CreatedBy: string;
  CreatedAt: Date;
  UpdatedBy: string;
  UpdatedAt: Date;
  DeletedAt?: Date;
  KeywordID: string;
  Keyword: KeywordFromServer;
  ParentKeywordID: string;
  ParentKeyword: KeywordFromServer;
}

export type KeywordFromServer = {
    ID: string;
    Status: string;
    CreatedBy: string;
    CreatedAt: Date
    UpdatedBy: string;
    UpdatedAt: Date;
    DeletedAt?: Date;
    Name: string;
    Code: string;
}


const RestHTTPServerAxios = async ():Promise<AxiosInstance> => {
  const e = await env()

  const instance = axios.create({
    baseURL: e.RestHTTPServer?.Url,
  });
  return instance
}

export const getCategories = async (): Promise<Result> => {
  const r: Result = { success: false }
  try {
    const axios = await RestHTTPServerAxios()
    const data = await axios.get('api/categories')
    const dataAry: CategoryFromServer[] = data.data?.Categories || []

    r.GetCategoriesResult = dataAry
    r.success = true
    return r
  } catch (e) {
    r.success = false
    r.message = String(e)
    r.error = e
    return r
  }

}

export const PutCategory = async (...rest: string[]): Promise<Result> => {
  const r: Result = { success: false }
  try {
    const axios = await RestHTTPServerAxios()
    const data = await axios.put('api/categories', {
      "Categories": rest,
    })
    console.log(data.data)

    r.success = true
    return r
  } catch (e) {
    r.success = false
    r.message = String(e)
    return r
  }

}