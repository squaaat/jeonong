import { createAction } from 'redux-actions'
import { Category } from 'store/models/Category'
import { RestAxios } from 'utils/rest'

export const initialState = {

}

export const API_GET_CATEGORIES = "category/API_GET_CATEGORIES"
export const apiGetCategories = createAction(
  API_GET_CATEGORIES,
  () => RestAxios().then(axios => axios.get('api/categories')),
)

export const API_PUT_CATEGORY = "category/API_PUT_CATEGORY"
export const apiPutCategory = createAction(
  API_PUT_CATEGORY,
  (c: Category) => RestAxios().then(axios => axios.put('api/categories', {
    Category: c
  })),
)
