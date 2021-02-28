import { RestAxios } from 'models/rest'

export type Category = {
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

  CreatedAt: Date;
  CreatedBy?: string;
  UpdatedAt: Date;
  UpdatedBy?: string;
  DeletedAt?: Date;
}

export const getCategories = async (): Promise<Category[]> => {
  const restAxios = await RestAxios()
  const res = await restAxios.get('api/categories')
  const data: Category[] = res.data?.Categories || []
  return data
}
