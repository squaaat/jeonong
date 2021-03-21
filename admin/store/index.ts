import { StoreEnhancer, createStore, compose, applyMiddleware } from 'redux';
import { MakeStore, createWrapper, Context, HYDRATE } from 'next-redux-wrapper';

import { composeWithDevTools } from 'redux-devtools-extension';

import ReduxThunk from 'redux-thunk'
import ReduxLogger from 'redux-logger'

import { rootReducer, RootState } from 'store/reducers'

const composeEnhancer: StoreEnhancer | any = process.env.NODE_ENV === 'production' ? compose : composeWithDevTools
export const configureStore = (preloadState: any) => {
  const middlewares: any = [
    ReduxThunk,
    ReduxLogger,
  ]; // 미들웨어들을 넣으면 된다.
  const enhancer = composeEnhancer(applyMiddleware(...middlewares))

  const store = createStore(rootReducer, preloadState, enhancer);
  return store;
}


// create a makeStore function
export const makeStore: MakeStore<RootState> = (_: Context) => configureStore(undefined);

// export an assembled wrapper
export const wrapper = createWrapper<RootState>(makeStore, {debug: true});