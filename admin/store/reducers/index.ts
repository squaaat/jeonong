import { combineReducers } from 'redux';

import { tick } from 'store/reducers/tick'

export const rootReducer = combineReducers({
  tick,
})

// 루트 리듀서의 반환값를 유추해줍니다
// 추후 이 타입을 컨테이너 컴포넌트에서 불러와서 사용해야 하므로 내보내줍니다.
export type RootState = ReturnType<typeof rootReducer>;
