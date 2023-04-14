export interface User {
  username?: string
  password: string
  email: string
  user_id?: string
  full_name?:string
  phone_number?:string
  current_conversations?:any[]
}
