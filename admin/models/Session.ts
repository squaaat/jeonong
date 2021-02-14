import { User, MockUser } from 'models/User'

export type Session = {
  user: User;
}


export const MockSession: Session = {
  user: MockUser,
}