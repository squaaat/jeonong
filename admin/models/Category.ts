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

  CreatedAt?: Date;
  CreatedBy?: string;
  UpdatedAt?: Date;
  UpdatedBy?: string;
  DeletedAt?: Date;
}

export const getCategories = async (): Promise<Category[]> => {
  const restAxios = await RestAxios()
  const res = await restAxios.get('api/categories')
  const data: Category[] = res.data?.Categories || []
  return data
}

export const putCategories = async (c: Category): Promise<Category> => {
  const restAxios = await RestAxios()
  const res = await restAxios.put('api/categories', {
    "Category": {
      "Name": "테스트7",
      "Code": "test7",
      "Depth": 1,
      "Sort": 2,
      "Category1ID": "",
      "Category2ID": "",
      "Category3ID": "",
      "Category4ID": ""
    }
  }
  )
  const data: Category = res.data?.Category || c
  return data

}
// {
//   "Category": {
//     "Name": "테스트3",
//     "Code": "test3",
//     "Depth": 1,
//     "Sort": 2,
//   }
// }

// {
//   "Category":{
//     "Sort":4,
//     "Name":"테스트4",
//     "Code":"test4",
//     "Depth":1
//   }
// }
