import { RestAxios } from 'utils/rest'

export interface ICategory {
  ID: string;
  Name: string;
  Code: string;
  FullName: string;
  Depth: number;
  Status: string;
  Sort: number;

  Category1ID?: string;
  Category2ID?: string;
  Category3ID?: string;
  Category4ID?: string;

  CreatedAt?: Date;
  CreatedBy?: string;
  UpdatedAt?: Date;
  UpdatedBy?: string;
  DeletedAt?: Date;
}
export class Category implements ICategory {
  ID = '';
  Name = '';
  Code = '';
  FullName = '';
  Depth = 1;
  Status = 'IDLE';
  Sort = 0;

  Category1ID = '';
  Category2ID = '';
  Category3ID = '';
  Category4ID = '';

  CreatedAt = new Date();
  CreatedBy = '';
  UpdatedAt = new Date();
  UpdatedBy = '';
  DeletedAt = new Date();
}


export const getCategories = async (): Promise<Category[]> => {
  const restAxios = await RestAxios()
  const res = await restAxios.get('api/categories')
  const data: Category[] = res.data?.Categories || []
  return data
}

export const putCategory = async (c: Category): Promise<Category> => {
  const restAxios = await RestAxios()
  const res = await restAxios.put('api/categories', {
    "Category": c,
  })
  const data: Category = res.data?.Category || c
  return data
}
