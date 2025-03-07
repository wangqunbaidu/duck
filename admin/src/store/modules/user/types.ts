export interface UserState {
  id: number;
  nickname?: string;
  username?: string;
  avatar?: string;
  gender?: string;
  phone?: string;
  email?: string;
  description?: string;
  status: number;
  roles?: string[];
  role: string; // TODO 应该是多个角色，不应该是一个角色
}
