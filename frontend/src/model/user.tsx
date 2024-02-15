export interface UserModel {
  user_id: string;
  display_name: string;
  user_name: string;
  class_id: string;
  mail_address: string;

  is_temp_user: boolean;
  is_admin: boolean;
  is_teacher: boolean;

  created_at: number;
  updated_at: number;
}
