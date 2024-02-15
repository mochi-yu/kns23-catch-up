import CheckNotRegister from "@/components/auth/check_not_register";
import { LogoutButton } from "@/components/common/button_logout";
import { Typography } from "@mui/material";

export default function RegisterPage() {
  return (
    <>
      <CheckNotRegister>
        <Typography>登録ページです</Typography>
        <Typography>対象のメールアドレスは{undefined}です。</Typography>
        <LogoutButton />
      </CheckNotRegister>
    </>
  );
}
