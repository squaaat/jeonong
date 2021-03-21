import { User, MockUser } from 'store/models/User'

export type Session = {
  user: User;
}


export const MockSession: Session = {
  user: MockUser,
}
