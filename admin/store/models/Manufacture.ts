import { fromJS, Record } from 'immutable'
import { RestAxios } from 'utils/rest'

export interface IManufacture {
  ID: string;
  Status: string;

  Name: string;
  Code: string;
  CompanyRegistrationNumber: string;

  CreatedAt?: Date;
  CreatedBy?: string;
  UpdatedAt?: Date;
  UpdatedBy?: string;
  DeletedAt?: Date;
}

export class Manufacture implements IManufacture{
  ID= ''
  Status= ''
  Name= ''
  Code= ''
  CompanyRegistrationNumber= ''

  CreatedAt= new Date()
  CreatedBy= ''
  UpdatedAt= new Date()
  UpdatedBy= ''
  DeletedAt= new Date()

  Record(): Record.Factory<Manufacture> {
    return Record(fromJS(this))
  }
}


export const getManufactures = async (): Promise<Manufacture[]> => {
  const restAxios = await RestAxios()
  const res = await restAxios.get('api/manufactures')
  const data: Manufacture[] = res.data?.Manufactures || []
  return data
}

export const putManufacture = async (c: Manufacture): Promise<Manufacture> => {
  const restAxios = await RestAxios()
  const res = await restAxios.put('api/manufactures', {
    "Manufacture": c,
  })
  const data: Manufacture = res.data?.Manufacture || c
  return data
}

