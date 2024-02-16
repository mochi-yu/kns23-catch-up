import CheckNotRegister from "@/components/auth/check_not_register";
import { LogoutButton } from "@/components/common/button/logout";
import { Typography } from "@mui/material";
import { RegisterForm } from "./register_form";

export default function RegisterPage() {
  return (
    <>
      <CheckNotRegister>
        <Typography>登録ページです</Typography>
        <RegisterForm />
        <LogoutButton />
      </CheckNotRegister>
    </>
  );
}
