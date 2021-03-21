import { HYDRATE } from 'next-redux-wrapper';
import { handleActions } from 'redux-actions'

// 이 리덕스 모듈에서 관리 할 상태의 타입을 선언합니다
export type TickState = {
  tick: string;
};

// 초기상태를 선언합니다.
export const tickInitState: TickState = {
  tick: 'init'
};

export const tick = handleActions({
  'TICK': (state, action: any) => ({ ...state, tick: action.payload || '' }),
  [HYDRATE]: (state, action: any) => ({ ...state, ...action.payload.tick }),
}, tickInitState)
