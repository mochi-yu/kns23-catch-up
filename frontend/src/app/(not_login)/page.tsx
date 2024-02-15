import CheckRootPage from "@/components/auth/check_root_page";
import { LogoutButton } from "@/components/common/button_logout";
import { GoogleLoginButton } from "@/components/common/button_login_google";

export default function RootPage() {
  return (
    <>
      <CheckRootPage>
        <GoogleLoginButton />
        <LogoutButton />
      </CheckRootPage>
    </>
  );
}
