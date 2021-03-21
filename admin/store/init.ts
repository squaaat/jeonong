import { Category } from 'store/models/Category'
import { Manufacture } from 'store/models/Manufacture'
import { Record } from 'immutable'

export interface IInitialState {
  Components: Record.Factory<IComponents>;
}

interface IComponents {
  CategoryManager: Record.Factory<Category>;
  ManufactureManager: Record.Factory<Manufacture>;
}

class Components implements IComponents {
  CategoryManager = Record(new Category());
  ManufactureManager = (new Manufacture()).Record();
}

export const initialState: IInitialState = {
  Components: Record(new Components()),
}
