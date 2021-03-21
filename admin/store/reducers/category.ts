import { handleActions } from 'redux-actions'
import * as actionsCategory from 'store/actions/category'


export default handleActions({
  [actionsCategory.API_GET_CATEGORIES]: state => state,
  [actionsCategory.API_PUT_CATEGORY]: state => state,
}, actionsCategory.initialState)